package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/type_game/components"
)

func main() {
    // instantiate
    width := int(os.Stdout.Fd()) // window width
	m := components.NewApp(width)

    // run
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
