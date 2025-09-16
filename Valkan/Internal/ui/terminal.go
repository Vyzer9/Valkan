package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Cores ANSI
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Bold   = "\033[1m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
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
 __   __  _______  ___      ___   _  _______  __    _
|  | |  ||   _   ||   |    |   | | ||   _   ||  |  | |
|  |_|  ||  |_|  ||   |    |   |_| ||  |_|  ||   |_| |
|       ||       ||   |    |      _||       ||       |
|       ||       ||   |___ |     |_ |       ||  _    |
 |     | |   _   ||       ||    _  ||   _   || | |   |
  |___|  |__| |__||_______||___| |_||__| |__||_|  |__|

`

func ShowBanner() {
	fmt.Print(Red, Bold)
	fmt.Println(valkanDragon)
	fmt.Println(valkanText)
	fmt.Print(Reset)

	printSystemInfo()
}

func printSystemInfo() {
	fmt.Println(Yellow + "───────────────────────────────────────────────" + Reset)
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

	fmt.Println(Yellow + "───────────────────────────────────────────────" + Reset)
}

func ShowMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println(Bold + "Menu:" + Reset)
		fmt.Println("1) Scanner")
		fmt.Println("2) Help")
		fmt.Println("3) Sair")
		fmt.Print("Escolha uma opção: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("\nDigite o IP ou host para escanear: ")
			host, _ := reader.ReadString('\n')
			host = strings.TrimSpace(host)

			if host == "" {
				fmt.Println(Red + "Host inválido, voltando ao menu." + Reset)
				continue
			}

			// Aqui você pode chamar seu scanner, exemplo:
			// scanner.StartScan(host)
			fmt.Printf("\nIniciando scan no host %s...\n", host)
			// Placeholder: simular scan
			fmt.Println("Scan finalizado.\n")

		case "2":
			showHelp()

		case "3":
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println(Red + "Opção inválida, tente novamente." + Reset)
		}
	}
}

func showHelp() {
	fmt.Println(`
` + Bold + `VALKAN - Network Recon Tool - Help` + Reset + `

1) Scanner
   - Escaneia portas TCP abertas no host alvo.
   - Informe IP ou domínio.
   - Busca portas abertas e tenta identificar serviços.

2) Help
   - Mostra este menu de ajuda.

3) Sair
   - Fecha o programa.

Dicas:
- Use IPs válidos (ex: 192.168.1.1) ou domínios (ex: example.com).
- Scan rápido: portas comuns (1-1024).
- Scan completo: todas as portas (1-65535), pode demorar.
- Consulte a documentação para mais informações.

`)
}
