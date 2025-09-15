package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Vyzer9/Valkan/Valkan/Internal/scanner"
)

// OutputFormat define o tipo de saída
type OutputFormat string

const (
	FormatJSON  OutputFormat = "json"
	FormatPlain OutputFormat = "txt"
	FormatTable OutputFormat = "table"
)

// Export escreve os resultados no formato desejado
func Export(results []scanner.PortScanResult, format OutputFormat, file string) error {
	switch format {
	case FormatJSON:
		return exportJSON(results, file)
	case FormatPlain:
		return exportPlain(results, file)
	case FormatTable:
		return exportTable(results)
	default:
		return fmt.Errorf("formato desconhecido: %s", format)
	}
}

func exportJSON(results []scanner.PortScanResult, file string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

func exportPlain(results []scanner.PortScanResult, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, r := range results {
		if r.Open {
			fmt.Fprintf(f, "Porta %d/%s aberta - Motivo: %s\n", r.Port, r.Protocol, r.Reason)
			if r.Banner != "" {
				fmt.Fprintf(f, " └─ Banner: %q\n", r.Banner)
			}
		}
	}
	return nil
}

func exportTable(results []scanner.PortScanResult) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "PORTA\tPROTOCOLO\tSTATUS\tMOTIVO\tBANNER")
	for _, r := range results {
		status := "fechada"
		if r.Open {
			status = "aberta"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%.30q\n", r.Port, r.Protocol, status, r.Reason, r.Banner)
	}
	return w.Flush()
}
