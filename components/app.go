package components

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	// estados de la aplicacion
	menu = iota
	inGame
	resumen
)

type App struct {
	Quit     bool
	Mode     int
	Menu     mainMenu
	Game     Typer
	appWidth int
}

func NewApp(size int) App {
	return App{
		Quit:     false,
		Mode:     menu,
		Menu:     NewMainMenu(),
		Game:     NewTyper(size),
		appWidth: size,
	}
}

// INFO: no se utiliza
func (m App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
    // cuando se encuentra inGame, delegar al typer
	if a.Mode == inGame {
		// si el juego termina reiniciar los valores y mostrar el menu principal
		if a.Game.Done {
			a.Game.Done = false
			a.Mode = menu
			return a, nil
		}
		// actualizar el juego
		a.Game, cmd = a.Game.Update(msg)
		return a, cmd
	}

	// Filtrar y administrar las acciones cuando una tecla es oprimida
	switch msg := msg.(type) {
	// handle when a key is pressed
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			println("See you latter !!!")
			time.Sleep(time.Second * 1)
			return a, tea.Quit
		}

		// TODO: revisar probablemente
		if msg.String() == "enter" {
			// ver cual opcion selecciono el usuario
			switch a.Menu.List.SelectedItem().FilterValue() {
			case "jugar": // crear una nueva instancia del typer
				a.Mode = inGame
				a.Game = NewTyper(a.appWidth)

			case "offline": // nuevo typer pero con el archivo local
				// TODO: implementar los archivos locales
				a.Mode = inGame
				a.Game = NewTyper(a.appWidth, "")

			case "cargar":
				// TODO: la cargar de nuevas palabras
				return a, tea.Quit
			}
		}
		// actualizar el menu
		a.Menu, cmd = a.Menu.Update(msg)

	// handle when the window is resized
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		a.Menu.List.SetSize(msg.Width-h, msg.Height-v)
		a.appWidth = msg.Width
	}
	return a, cmd
}

func (m App) View() string {
	// si el juego continua entonces mostrar el juego
	if m.Mode == inGame {
		return docStyle.Render(m.Game.View())
	}
	// mostrar el menu principa
	return docStyle.Render(m.Menu.View())
}
