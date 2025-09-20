package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/Vyzer9/Valkan/Valkan/Internal/detection"
	"github.com/Vyzer9/Valkan/Valkan/Internal/discovery"
	"github.com/Vyzer9/Valkan/Valkan/Internal/plugins"
	"github.com/Vyzer9/Valkan/Valkan/Internal/recon" // import do recon adicionado
	"github.com/Vyzer9/Valkan/Valkan/Internal/scanner"
)

// Cores ANSI
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Bold   = "\033[1m"
)

// Dragão em vermelho, mais "sério"
const valkanDragon = `
              ___====-_  _-====___
        _--^^^#####//      \\#####^^^--_
      _-^##########// (    ) \\##########^-_
     -############//  |\^^/|  \\############-
   _/############//   (@::@)   \\############\_
  /#############((     \\//     ))#############\
 -###############\\    (oo)    //###############-
-#################\\  / VV \  //#################-
-###################\\/      \//###################-
_#/|##########/\######(   /\   )######/\##########|\#_
|/ |#/\#/\#/\/  \#/\##\  ||  /##/\#/  \/\#/\#/\#| \|
` + "`" + `  |/  V  V  ` + "`" + `   V  \\#\| | | | | |/#/  V   '  V  V  \|  
` + "`" + `   ` + "`" + `   ` + "`" + `   ` + "`" + `   / | | | | | | \  
               (  | | | | | |  )  
              __\ | | | | | | /__  
             (vvv(VVV)(VVV)vvv)  
`

const valkanText = `
██╗   ██╗ █████╗ ██╗     ██╗  ██╗ █████╗ ███╗   ██╗
██║   ██║██╔══██╗██║     ██║ ██╔╝██╔══██╗████╗  ██║
██║   ██║███████║██║     █████╔╝ ███████║██╔██╗ ██║
╚██╗ ██╔╝██╔══██║██║     ██╔═██╗ ██╔══██║██║╚██╗██║
 ╚████╔╝ ██║  ██║███████╗██║  ██╗██║  ██║██║ ╚████║
  ╚═══╝  ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝

`

func ShowBanner() {
	fmt.Print(Red, Bold)
	fmt.Println(valkanDragon)
	fmt.Println(valkanText)
	fmt.Print(Reset)

	printSystemInfo()
}

func printSystemInfo() {
	fmt.Println("───────────────────────────────────────────────")
	fmt.Printf(" Sistema: %s\n", runtime.GOOS)
	fmt.Printf(" Arquitetura: %s\n", runtime.GOARCH)

	host, err := exec.Command("hostname").Output()
	if err == nil {
		fmt.Printf(" Hostname: %s\n", strings.TrimSpace(string(host)))
	}

	if runtime.GOOS == "linux" {
		uname, err := exec.Command("uname", "-r").Output()
		if err == nil {
			fmt.Printf(" Kernel: %s\n", strings.TrimSpace(string(uname)))
		}
	}

	fmt.Println("───────────────────────────────────────────────")
}

func ShowMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println(Yellow + Bold + "Menu:" + Reset)
		fmt.Println(Yellow + "1) Scanner" + Reset)
		fmt.Println(Yellow + "2) Discovery" + Reset)
		fmt.Println(Yellow + "3) Recon (Subdomain Finder)" + Reset)
		fmt.Println(Yellow + "4) Help" + Reset)
		fmt.Println(Yellow + "5) Sair" + Reset)
		fmt.Print(Reset + "Escolha uma opção: " + Reset)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			// Escolher protocolo TCP/UDP
			fmt.Println()
			fmt.Println(Yellow + Bold + "Escolha o protocolo:" + Reset)
			fmt.Println(Yellow + "1) TCP" + Reset)
			fmt.Println(Yellow + "2) UDP" + Reset)
			fmt.Print(Reset + "Protocolo: " + Reset)

			protoOption, _ := reader.ReadString('\n')
			protoOption = strings.TrimSpace(protoOption)

			var protocol string
			switch protoOption {
			case "1":
				protocol = "tcp"
			case "2":
				protocol = "udp"
			default:
				fmt.Println(Blue + "Protocolo inválido, voltando ao menu." + Reset)
				continue
			}

			// Escolher tipo de scan (portas)
			fmt.Println()
			fmt.Println(Yellow + Bold + "Escolha o tipo de scan:" + Reset)
			fmt.Println(Yellow + "1) Scan rápido (portas 1-1024)" + Reset)
			fmt.Println(Yellow + "2) Scan completo (portas 1-65535)" + Reset)
			fmt.Print(Reset + "Opção: " + Reset)

			scanOption, _ := reader.ReadString('\n')
			scanOption = strings.TrimSpace(scanOption)

			var startPort, endPort int
			switch scanOption {
			case "1":
				startPort, endPort = 1, 1024
			case "2":
				startPort, endPort = 1, 65535
			default:
				fmt.Println(Blue + "Opção inválida, voltando ao menu." + Reset)
				continue
			}

			// Solicitar host/IP
			fmt.Print("\nDigite o IP ou host para escanear: ")
			host, _ := reader.ReadString('\n')
			host = strings.TrimSpace(host)

			if host == "" {
				fmt.Println(Blue + "Host inválido, voltando ao menu." + Reset)
				continue
			}

			fmt.Printf("\nIniciando scan no host %s, portas %d-%d, protocolo %s...\n", host, startPort, endPort, strings.ToUpper(protocol))

			ctx := context.Background()
			timeout := 2 * time.Second
			concurrency := 100

			results := scanner.ScanRangeConcurrent(ctx, host, startPort, endPort, timeout, protocol, concurrency)

			fmt.Println()
			fmt.Println(Blue + Bold + "Resultados do Scan:" + Reset)

			openCount := 0
			for _, res := range results {
				if res.Open {
					openCount++
					service := detection.DetectService(res.Banner)
					fmt.Printf(Green+"Porta %d aberta (%s) - Serviço detectado: %s\n"+Reset, res.Port, res.Protocol, service)

					if service == "HTTP" && protocol == "tcp" {
						banner, err := plugins.GrabHTTPBanner(fmt.Sprintf("%s:%d", host, res.Port), timeout)
						if err == nil {
							fmt.Printf(Purple+"Banner HTTP detalhado:\n%s\n"+Reset, banner)
						}
					}
				}
			}

			if openCount == 0 {
				fmt.Println(Yellow + "Nenhuma porta aberta encontrada." + Reset)
			}

			fmt.Println("Scan finalizado.\n")

			// Salvar resultados em arquivo
			filename := "resultados_scan.txt"
			err := saveResultsToFile(results, filename)
			if err != nil {
				fmt.Println(Red, "Erro ao salvar resultados:", err, Reset)
			} else {
				fmt.Println(Green, "Resultados salvos em", filename, Reset)
			}

		case "2":
			fmt.Println(Yellow + "Iniciando Discovery..." + Reset)

			reader := bufio.NewReader(os.Stdin)

			// Pedir CIDR para o usuário
			fmt.Print("Digite a rede no formato CIDR (ex: 192.168.1.0/24): ")
			cidr, _ := reader.ReadString('\n')
			cidr = strings.TrimSpace(cidr)
			if cidr == "" {
				fmt.Println(Red + "CIDR inválido." + Reset)
				continue
			}

			// Pedir método
			fmt.Print("Digite o método (icmp/tcp): ")
			method, _ := reader.ReadString('\n')
			method = strings.TrimSpace(method)
			if method != "icmp" && method != "tcp" {
				fmt.Println(Red + "Método inválido." + Reset)
				continue
			}

			// Se método for TCP, pedir porta
			tcpPort := 80 // valor padrão
			if method == "tcp" {
				fmt.Print("Digite a porta TCP para testar (ex: 80): ")
				portInput, _ := reader.ReadString('\n')
				portInput = strings.TrimSpace(portInput)
				_, err := fmt.Sscanf(portInput, "%d", &tcpPort)
				if err != nil || tcpPort <= 0 || tcpPort > 65535 {
					fmt.Println(Red + "Porta inválida." + Reset)
					continue
				}
			}

			ctx := context.Background()
			timeout := 2 * time.Second

			results, err := discovery.RunDiscovery(ctx, cidr, method, timeout, tcpPort)
			if err != nil {
				fmt.Println(Red + "Erro ao executar discovery: " + err.Error() + Reset)
			} else {
				if len(results) == 0 {
					fmt.Println(Yellow + "Nenhum host ativo encontrado." + Reset)
				} else {
					fmt.Println(Green + "Resultados do Discovery:" + Reset)
					for _, res := range results {
						fmt.Printf("IP: %s, Vivo: %t, Método: %s\n", res.IP, res.Alive, res.Method)
					}
				}
			}

		case "3":
			fmt.Println(Yellow + "Iniciando Recon (Busca de Subdomínios)..." + Reset)

			fmt.Print("Digite o domínio (ex: exemplo.com): ")
			domain, _ := reader.ReadString('\n')
			domain = strings.TrimSpace(domain)
			if domain == "" {
				fmt.Println(Red + "Domínio inválido." + Reset)
				continue
			}

			fmt.Print("Digite o timeout em segundos (ex: 2): ")
			timeoutStr, _ := reader.ReadString('\n')
			timeoutStr = strings.TrimSpace(timeoutStr)
			timeoutSec, err := time.ParseDuration(timeoutStr + "s")
			if err != nil || timeoutSec <= 0 {
				fmt.Println(Red + "Timeout inválido. Usando padrão 2 segundos." + Reset)
				timeoutSec = 2 * time.Second
			}

			fmt.Print("Digite a concorrência (número de goroutines, ex: 10): ")
			concurrencyStr, _ := reader.ReadString('\n')
			concurrencyStr = strings.TrimSpace(concurrencyStr)
			concurrency := 10
			if concurrencyStr != "" {
				fmt.Sscanf(concurrencyStr, "%d", &concurrency)
				if concurrency <= 0 {
					concurrency = 10
				}
			}

			results := recon.FindSubdomains(domain, timeoutSec, concurrency)
			if len(results) == 0 {
				fmt.Println(Red + "Nenhum subdomínio encontrado ativo." + Reset)
			} else {
				fmt.Println(Green + "Subdomínios encontrados:" + Reset)
				for _, res := range results {
					fmt.Printf("Subdomínio: %s | IP: %s\n", res.Subdomain, res.IP)
				}
			}

		default:
			fmt.Println(Red + "Opção inválida, tente novamente." + Reset)
		}

		case "4":
			showHelp()

		case "5":
			fmt.Println("Saindo...")
			return
	}
}

func saveResultsToFile(results []scanner.PortScanResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, res := range results {
		status := "fechada"
		if res.Open {
			status = "aberta"
		}
		line := fmt.Sprintf("Porta %d (%s): %s - %s\n", res.Port, res.Protocol, status, res.Reason)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func showHelp() {
	fmt.Println(`
` + Bold + `VALKAN - Network Recon Tool - Help` + Reset + `

1) Scanner
   - Escaneia portas TCP ou UDP abertas no host alvo.
   - Informe IP ou domínio.
   - Escolha protocolo TCP ou UDP.
   - Busca portas abertas e tenta identificar serviços (apenas TCP).
   - Salva resultados em arquivo "resultados_scan.txt".

2) Discovery
   - Executa uma varredura de descoberta na rede (ex: ping sweep, hosts ativos).

3) Help
   - Mostra este menu de ajuda.

4) Sair
   - Fecha o programa.

5) Recon (Subdomain Finder)
   - Realiza busca de subdomínios ativos para um domínio informado.
   - Permite definir timeout e concorrência.

Dicas:
- Use IPs válidos (ex: 192.168.1.1) ou domínios (ex: example.com).
- Scan rápido: portas comuns (1-1024).
- Scan completo: todas as portas (1-65535), pode demorar.
- UDP scan não identifica serviços nem banner.
- Consulte a documentação para mais informações.

`)
}
