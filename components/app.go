package components

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// estados en los que se puede encontrar el juego
const (
	menu = iota
	game
	resumen
)

type App struct {
	Quit    bool
	Mode    int
	Menu    mainMenu
	Playing bool
	Game    Typer
}

func NewApp() App {
	return App{
		Quit:    false,
		Mode:    menu,
		Menu:    NewMainMenu(),
		Playing: false,
		Game:    NewTyper(200),
	}
}

// INFO: no se utiliza
func (m App) Init() tea.Cmd {
	return nil
}

func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// si el juego termina reiniciar los valores y mostrar el menu principal
	if m.Playing && m.Game.Done {
		m.Playing = false
		m.Game.Done = false
		return m, nil
	}

	// actualizar el juego si es que se esta jugando
	var cmd tea.Cmd
	if m.Playing {
		m.Game, cmd = m.Game.Update(msg)
		return m, cmd
	}

	// Filtrar y administrar las acciones cuando una tecla es oprimida
	switch msg := msg.(type) {
	// handle when a key is pressed
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			println("See you latter !!!")
			time.Sleep(time.Second * 1)
			return m, tea.Quit
		}

		// TODO: probablemente pasar a iota con consts
		if msg.String() == "enter" {
			switch m.Menu.List.SelectedItem().FilterValue() {
			case "jugar":
				m.Playing = !m.Playing
				m.Game = NewTyper(200)

			case "offline":
				m.Playing = !m.Playing

			case "cargar":
				return m, tea.Quit
			}
		}
		m.Menu, cmd = m.Menu.Update(msg)

	// handle when the window is resized
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.Menu.List.SetSize(msg.Width-h, msg.Height-v)
		m.Game.OutputSize = msg.Width - h
	}
	return m, cmd
}

func (m App) View() string {
	// si el juego continua entonces mostrar el juego
	if m.Playing {
		return docStyle.Render(m.Game.View())
	}
	// mostrar el menu principa
	return docStyle.Render(m.Menu.View())
}
