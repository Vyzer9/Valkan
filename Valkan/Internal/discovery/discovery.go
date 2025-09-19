package discovery

import (
	"context"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// DiscoveryResult armazena o resultado da descoberta de um host
type DiscoveryResult struct {
	IP     string
	Alive  bool
	Method string
}

// RunDiscovery executa descoberta em uma rede no formato CIDR usando método "icmp" ou "tcp"
// Agora aceita portaTCP como parâmetro, caso o método seja "tcp"
func RunDiscovery(ctx context.Context, cidr string, method string, timeout time.Duration, tcpPort int) ([]DiscoveryResult, error) {
	ips, err := ExpandCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("CIDR inválido: %v", err)
	}

	results := DiscoverHostsConcurrently(ctx, ips, method, timeout, tcpPort)
	return results, nil
}

// PingICMP envia um pacote ICMP Echo Request para um IP
func PingICMP(ip string, timeout time.Duration) bool {
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return false // Sem permissão para socket raw
	}
	defer c.Close()

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   1234,
			Seq:  1,
			Data: []byte("HELLO-PING"),
		},
	}
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		return false
	}

	dst, err := net.ResolveIPAddr("ip4", ip)
	if err != nil {
		return false
	}

	if _, err := c.WriteTo(msgBytes, dst); err != nil {
		return false
	}

	c.SetReadDeadline(time.Now().Add(timeout))

	buf := make([]byte, 1500)
	n, _, err := c.ReadFrom(buf)
	if err != nil {
		return false
	}

	rm, err := icmp.ParseMessage(1, buf[:n])
	if err != nil {
		return false
	}

	return rm.Type == ipv4.ICMPTypeEchoReply
}

// PingTCP tenta conectar em uma porta TCP
func PingTCP(ip string, port int, timeout time.Duration) bool {
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// ExpandCIDR retorna todos os IPs dentro de um CIDR
func ExpandCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ipCopy := make(net.IP, len(ip))
		copy(ipCopy, ip)
		ips = append(ips, ipCopy.String())
	}

	// Remove endereço de rede e broadcast
	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}
	return ips, nil
}

// incIP incrementa o IP
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// DiscoverHostsConcurrently faz a varredura concorrente
func DiscoverHostsConcurrently(ctx context.Context, ips []string, method string, timeout time.Duration, tcpPort int) []DiscoveryResult {
	results := make([]DiscoveryResult, 0, len(ips))
	resultsChan := make(chan DiscoveryResult, len(ips))

	for _, ip := range ips {
		go func(ip string) {
			alive := false

			if method == "icmp" {
				alive = PingICMP(ip, timeout)
			} else if method == "tcp" {
				alive = PingTCP(ip, tcpPort, timeout)
			}

			resultsChan <- DiscoveryResult{
				IP:     ip,
				Alive:  alive,
				Method: method,
			}
		}(ip)
	}

	for i := 0; i < len(ips); i++ {
		select {
		case res := <-resultsChan:
			if res.Alive {
				results = append(results, res)
			}
		case <-ctx.Done():
			return results
		}
	}

	return results
}
