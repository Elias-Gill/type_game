package components

import "github.com/charmbracelet/lipgloss"

var (
	// docStyle corresponde al estilo utilizado para renderizar la app entera
	docStyle = lipgloss.NewStyle().Width(200).
			Height(20).
			Margin(1, 2)

	// styles to colorate strings while the user is typing
	goodStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#43d11c"))

	badStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#f40045"))

	doneStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ae67f0"))

	authorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#999999"))
)
