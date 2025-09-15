package scanner

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

type PortScanResult struct {
	Port     int
	Open     bool
	Protocol string
	Reason   string
	Banner   string // Novo campo para banner
}

func ScanPort(ctx context.Context, host string, port int, timeout time.Duration, protocol string) PortScanResult {
	address := fmt.Sprintf("%s:%d", host, port)

	switch protocol {
	case "udp":
		return scanUDP(ctx, address, port, timeout)
	case "tcp":
		return scanTCP(ctx, address, port, timeout)
	default:
		return PortScanResult{
			Port:     port,
			Open:     false,
			Protocol: protocol,
			Reason:   "protocolo não suportado",
			Banner:   "",
		}
	}
}

func scanTCP(ctx context.Context, address string, port int, timeout time.Duration) PortScanResult {
	dialer := net.Dialer{Timeout: timeout}
	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return PortScanResult{
			Port:     port,
			Open:     false,
			Protocol: "tcp",
			Reason:   err.Error(),
			Banner:   "",
		}
	}
	defer conn.Close()

	// Tenta enviar uma requisição HTTP genérica para obter banner
	_ = conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	_, _ = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))

	// Ler o possível banner da conexão
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	var banner string
	if err == nil && n > 0 {
		banner = string(buf[:n])
	}

	return PortScanResult{
		Port:     port,
		Open:     true,
		Protocol: "tcp",
		Reason:   "conectado",
		Banner:   banner,
	}
}

func scanUDP(ctx context.Context, address string, port int, timeout time.Duration) PortScanResult {
	var lastErr error
	retries := 2

	for i := 0; i < retries; i++ {
		conn, err := net.DialTimeout("udp", address, timeout)
		if err != nil {
			lastErr = err
			continue
		}

		_, err = conn.Write([]byte{0})
		if err != nil {
			lastErr = err
			conn.Close()
			continue
		}

		conn.SetReadDeadline(time.Now().Add(timeout))
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		conn.Close()

		if err == nil {
			return PortScanResult{
				Port:     port,
				Open:     true,
				Protocol: "udp",
				Reason:   "resposta recebida",
				Banner:   "",
			}
		} else if ne, ok := err.(net.Error); ok && ne.Timeout() {
			lastErr = err
			continue
		} else {
			lastErr = err
			break
		}
	}

	return PortScanResult{
		Port:     port,
		Open:     false,
		Protocol: "udp",
		Reason:   fmt.Sprintf("sem resposta (último erro: %v)", lastErr),
		Banner:   "",
	}
}

func ScanFullRange(ctx context.Context, host string, timeout time.Duration, protocol string, concurrency int) []PortScanResult {
	return ScanRangeConcurrent(ctx, host, 1, 65535, timeout, protocol, concurrency)
}

func ScanRangeConcurrent(ctx context.Context, host string, start, end int, timeout time.Duration, protocol string, concurrency int) []PortScanResult {
	var results []PortScanResult
	var mu sync.Mutex
	var wg sync.WaitGroup

	if concurrency <= 0 {
		concurrency = 100
	}

	ports := make(chan int, concurrency)
	var processedCount, openCount int
	startTime := time.Now()
	doneChan := make(chan struct{})

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range ports {
				select {
				case <-ctx.Done():
					return
				default:
				}

				res := ScanPort(ctx, host, port, timeout, protocol)
				mu.Lock()
				results = append(results, res)
				processedCount++
				if res.Open {
					openCount++
				}
				mu.Unlock()
			}
		}()
	}

	go func() {
		for port := start; port <= end; port++ {
			select {
			case <-ctx.Done():
				break
			case ports <- port:
			}
		}
		close(ports)
	}()

	go func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("\nScan cancelado!")
				return
			case <-doneChan:
				return
			case <-ticker.C:
				mu.Lock()
				done := processedCount
				opens := openCount
				mu.Unlock()

				percent := float64(done) / float64(end-start+1) * 100
				elapsed := time.Since(startTime)
				var eta time.Duration
				if done > 0 {
					eta = time.Duration(float64(elapsed) / float64(done) * float64(end-start+1-done))
				}

				fmt.Printf("\rEscaneando... %.2f%% concluído | Portas abertas: %d | Tempo decorrido: %s | ETA: %s",
					percent, opens, elapsed.Truncate(time.Second), eta.Truncate(time.Second))
			}
		}
	}()

	wg.Wait()
	close(doneChan)

	fmt.Println("\nScan concluído!")
	return results
}
