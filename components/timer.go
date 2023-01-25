package components

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: arreglar el timer
type Timer struct {
	stopwatch stopwatch.Model
	help      help.Model
	quitting  bool
}

func (m Timer) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m Timer) View() string {
	// Note: you could further customize the time output by getting the
	// duration from m.stopwatch.Elapsed(), which returns a time.Duration, and
	// skip m.stopwatch.View() altogether.
	s := m.stopwatch.View() + "\n"
	if !m.quitting {
		s = "Elapsed: " + s
	}
	return s
}

func (m Timer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, cmd
}
