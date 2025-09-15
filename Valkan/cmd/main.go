package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Vyzer9/Valkan/Valkan/Internal/scanner"
	"github.com/spf13/cobra"
)

func main() {
	var host string
	var ports string
	var timeout int

	var rootCmd = &cobra.Command{
		Use:   "valkan",
		Short: "Valkan - scanner de rede e exploração",
		Long:  "Valkan é uma ferramenta para scan de portas, exploração de vulnerabilidades e mais.",
	}

	var scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "Faz scan de portas em um host",
		Run: func(cmd *cobra.Command, args []string) {
			portRange := strings.Split(ports, "-")
			start, _ := strconv.Atoi(portRange[0])
			end, _ := strconv.Atoi(portRange[1])

			openPorts := scanner.ScanRange(host, start, end, time.Duration(timeout)*time.Millisecond)
			fmt.Printf("Portas abertas em %s: %v\n", host, openPorts)
		},
	}

	scanCmd.Flags().StringVarP(&host, "host", "H", "", "Host alvo (obrigatório)")
	scanCmd.Flags().StringVarP(&ports, "ports", "p", "1-1024", "Faixa de portas para scan (ex: 1-1024)")
	scanCmd.Flags().IntVarP(&timeout, "timeout", "t", 500, "Timeout em milissegundos para cada porta")
	scanCmd.MarkFlagRequired("host")

	rootCmd.AddCommand(scanCmd)
	rootCmd.Execute()
}
