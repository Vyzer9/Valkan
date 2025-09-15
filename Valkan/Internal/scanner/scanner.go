package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPort tenta se conectar na porta especificada do host.
// Retorna true se a conexão for bem sucedida (porta aberta).
func ScanPort(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// ScanRange faz scan das portas de start até end no host.
// Retorna um slice com as portas abertas.
func ScanRange(host string, start, end int, timeout time.Duration) []int {
	var openPorts []int
	for port := start; port <= end; port++ {
		if ScanPort(host, port, timeout) {
			openPorts = append(openPorts, port)
		}
	}
	return openPorts
}
