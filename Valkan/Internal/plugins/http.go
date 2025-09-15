package plugins

import (
	"net"
	"time"
)

// GrabHTTPBanner envia uma requisição HTTP e tenta capturar o banner.
func GrabHTTPBanner(addr string, timeout time.Duration) (string, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(timeout))
	req := "HEAD / HTTP/1.0\r\n\r\n"
	_, err = conn.Write([]byte(req))
	if err != nil {
		return "", err
	}

	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}
