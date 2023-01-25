package components

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewMainMenu() *Model {
	items := []list.Item{
		Item{Tit: "Jugar", Desc: "Citas famosas sacadas aleatoriamente de internet"},
		Item{Tit: "Jugar offline", Desc: "Citas sacadas del banco de citas local"},
		Item{Tit: "Personalizado", Desc: "Escribe tu propia cita personalizada y guardala en el banco de citas local"},
	}

	m := Model{List: list.New(items, list.NewDefaultDelegate(), 0, 0), Playing: false, Game: NewTyper()}
	m.List.Title = "Menu Principal"
	return &m
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Item struct {
	Tit, Desc string
}

func (i Item) Title() string       { return i.Tit }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Tit }

type Model struct {
	List    list.Model
	Playing bool
	Game    *Typer
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			println("see you latter")
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			if !m.Playing {
				m.Playing = true
			}
			return m, nil
		}

		// resize of the window
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	if m.Playing {
		return m.Game.Update(msg)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
    // si el juego acaba de empezar
	if !m.Playing {
		return docStyle.Render(m.List.View())
	}
    // si ya se termino de jugar
    if m.Playing && m.Game.Done {
        m.Playing = false
        m.Game.Done = false
		return docStyle.Render(m.List.View())
    }
	return m.Game.View()
}
