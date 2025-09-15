package scanner

import (
	"fmt"
	"sync"
	"time"
)

// Status armazena dados do progresso do scan
type Status struct {
	mu           sync.Mutex
	totalPorts   int
	checkedPorts int
	openPorts    int
	startTime    time.Time
}

// NewStatus cria um novo status com total de portas
func NewStatus(totalPorts int) *Status {
	return &Status{
		totalPorts: totalPorts,
		startTime:  time.Now(),
	}
}

// IncrementChecked incrementa a quantidade de portas verificadas
func (s *Status) IncrementChecked() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.checkedPorts++
}

// IncrementOpen incrementa a quantidade de portas abertas encontradas
func (s *Status) IncrementOpen() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.openPorts++
}

// Print atualiza o status no terminal
func (s *Status) Print() {
	s.mu.Lock()
	defer s.mu.Unlock()

	percent := float64(s.checkedPorts) / float64(s.totalPorts) * 100
	elapsed := time.Since(s.startTime).Round(time.Second)

	fmt.Printf("\rScan progress: %.2f%% | Ports checked: %d/%d | Open ports: %d | Elapsed time: %s",
		percent, s.checkedPorts, s.totalPorts, s.openPorts, elapsed)
}
