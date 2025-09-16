package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Vyzer9/Valkan/Valkan/Internal/scanner"
	"github.com/Vyzer9/Valkan/Valkan/Internal/ui"
	"github.com/spf13/cobra"
)

func main() {
	ui.ShowBanner()
	ui.ShowMenu()
	var host string
	var ports string
	var timeout int
	var full bool
	var protocol string
	var concurrency int

	// Comando raiz
	var rootCmd = &cobra.Command{
		Use:   "valkan",
		Short: "Valkan - scanner de rede e exploração",
		Long:  "Valkan é uma ferramenta para scan de portas, exploração de vulnerabilidades e mais.",
	}

	// Comando de scan
	var scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "Faz scan de portas em um host",
		Run: func(cmd *cobra.Command, args []string) {
			// Exibir banner com info do sistema
			ui.ShowBanner()

			startTime := time.Now()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var results []scanner.PortScanResult

			if full {
				fmt.Println("[*] Iniciando scan completo de 1 a 65535...")
				results = scanner.ScanFullRange(ctx, host, time.Duration(timeout)*time.Millisecond, protocol, concurrency)
			} else {
				portRange := strings.Split(ports, "-")
				if len(portRange) != 2 {
					fmt.Println("Faixa de portas inválida. Use o formato correto, ex: 1-1024")
					return
				}
				start, err1 := strconv.Atoi(portRange[0])
				end, err2 := strconv.Atoi(portRange[1])
				if err1 != nil || err2 != nil || start < 1 || end > 65535 || start > end {
					fmt.Println("Faixa de portas inválida. Use o formato correto, ex: 1-1024")
					return
				}
				fmt.Printf("[*] Iniciando scan de %d a %d...\n", start, end)
				results = scanner.ScanRangeConcurrent(ctx, host, start, end, time.Duration(timeout)*time.Millisecond, protocol, concurrency)
			}

			fmt.Printf("\n[+] Portas abertas em %s:\n", host)
			openCount := 0
			for _, r := range results {
				if r.Open {
					fmt.Printf("  ▸ Porta %d/%s aberta - Motivo: %s\n", r.Port, r.Protocol, r.Reason)
					if r.Banner != "" {
						fmt.Printf("    └─ Banner: %q\n", strings.TrimSpace(r.Banner))
					}
					openCount++
				}
			}

			if openCount == 0 {
				fmt.Println("  Nenhuma porta aberta encontrada.")
			}

			fmt.Printf("[*] Scan concluído em %s\n", time.Since(startTime))
		},
	}

	// Flags do comando scan
	scanCmd.Flags().StringVarP(&host, "host", "H", "", "Host alvo (obrigatório)")
	scanCmd.Flags().StringVarP(&ports, "ports", "p", "1-1024", "Faixa de portas para scan (ex: 1-1024)")
	scanCmd.Flags().IntVarP(&timeout, "timeout", "t", 500, "Timeout em milissegundos para cada porta")
	scanCmd.Flags().BoolVarP(&full, "full", "f", false, "Faz scan completo das 65535 portas")
	scanCmd.Flags().StringVarP(&protocol, "protocol", "P", "tcp", "Protocolo a ser escaneado (tcp ou udp)")
	scanCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 100, "Número de goroutines concorrentes para scan")
	scanCmd.MarkFlagRequired("host")

	// Adiciona comando ao root
	rootCmd.AddCommand(scanCmd)

	// Executa CLI
	rootCmd.Execute()
}
