package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const defaultLogDir = "logs"

func SaveScanResults(filename string, content string) error {
	if filename == "" {
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		filename = fmt.Sprintf("scan_%s.txt", timestamp)
	}

	// Garante que o diretório exista
	if err := os.MkdirAll(defaultLogDir, os.ModePerm); err != nil {
		return fmt.Errorf("erro ao criar diretório de log: %v", err)
	}

	fullPath := filepath.Join(defaultLogDir, filename)

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de log: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo de log: %v", err)
	}

	return nil
}
