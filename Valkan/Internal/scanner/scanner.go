package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPort tenta abrir conexão TCP na porta especificada
func ScanPort(protocol, hostname string, port int) bool {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, address, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
