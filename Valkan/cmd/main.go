package main

import (
	"fmt"

	"github.com/Vyzer9/Valkan/Valkan/Internal/scanner"
)

func main() {
	host := "scanme.nmap.org"
	port := 80

	if scanner.ScanPort("tcp", host, port) {
		fmt.Printf("Porta %d aberta em %s\n", port, host)
	} else {
		fmt.Printf("Porta %d fechada em %s\n", port, host)
	}
}
