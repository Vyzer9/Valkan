package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Vyzer9/Valkan/Valkan/Internal/scanner"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Título
	title := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("🐉 [green::b]VALKAN - Network TUI Scanner")

	// Cria o menu principal
	menu := tview.NewList().
		AddItem("Scan rápido (1-1024)", "Executa scan nas portas mais comuns", '1', func() {
			go executeScan(app, "127.0.0.1", 1, 1024)
		}).
		AddItem("Scan completo (1-65535)", "Scan total de todas as portas", '2', func() {
			go executeScan(app, "127.0.0.1", 1, 65535)
		}).
		AddItem("Ver resultados anteriores", "Mostra scans salvos em txt/json", '3', func() {
			showComingSoon(app)
		}).
		AddItem("Sair", "Sai do aplicativo", '4', func() {
			app.Stop()
		})

	menu.SetBorder(true).SetTitle(" Menu ").SetTitleAlign(tview.AlignLeft)

	// Layout principal
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(title, 3, 1, false).
		AddItem(menu, 0, 1, true)

	// Executa o app
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

// Função que executa o scan e mostra os resultados no TUI
func executeScan(app *tview.Application, host string, startPort, endPort int) {
	view := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetTextAlign(tview.AlignLeft).
		SetChangedFunc(func() {
			app.Draw()
		})

	view.SetBorder(true).SetTitle(" Resultados do Scan ")

	fmt.Fprintf(view, "[yellow]Escaneando %s de %d a %d...\n\n", host, startPort, endPort)

	app.QueueUpdateDraw(func() {
		app.SetRoot(view, true)
	})

	startTime := time.Now()

	results := scanner.ScanRangeConcurrent(context.Background(), host, startPort, endPort, 500*time.Millisecond, "tcp", 100)

	open := 0
	for _, r := range results {
		if r.Open {
			open++
			fmt.Fprintf(view, "[green]Porta %d/%s aberta - Motivo: %s\n", r.Port, r.Protocol, r.Reason)
			if r.Banner != "" {
				fmt.Fprintf(view, "  [gray]Banner: %s\n", strings.TrimSpace(r.Banner))
			}
		}
	}

	if open == 0 {
		fmt.Fprintf(view, "\n[red]Nenhuma porta aberta encontrada.")
	}

	fmt.Fprintf(view, "\n\n[white]Scan concluído em: [cyan]%s", time.Since(startTime))
	fmt.Fprintf(view, "\n[white]Pressione [::b]Ctrl+C[::-] para sair")

	view.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyRune && (key == 'q' || key == 'Q') {
			app.QueueUpdateDraw(func() {
				main()
			})
		}
	})
}

// Placeholder para a opção de "ver resultados anteriores"
func showComingSoon(app *tview.Application) {
	text := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("\n\n[green]🚧 Em breve: visualização de resultados anteriores!\n\nPressione [::b]Q[::-] para voltar.")

	text.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyRune && (key == 'q' || key == 'Q') {
			app.QueueUpdateDraw(func() {
				main()
			})
		}
	})

	app.SetRoot(text, true)
}
