package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylan0804/godocai/shared"
	"github.com/dylan0804/godocai/ui"
)

func main() {
	p := tea.NewProgram(
		ui.NewAppModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	shared.Program = p

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
