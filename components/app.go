package components

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	// estados de la aplicacion
	menu = iota
	inGame
	resumen
	local
)

type App struct {
	Quit        bool
	Mode        int
	Menu        MainMenu
	LocalQuotes QuotesMenu
	Game        Typer
	appWith     int
	appHeight   int
}

func NewApp() App {
	return App{
		Quit:        false,
		Mode:        menu,
		Menu:        NewMainMenu(),
		Game:        Typer{},
		LocalQuotes: QuotesMenu{},
	}
}

func (m App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// actualizar el tamano de la app
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.appWith = msg.Width
		a.appWith = msg.Height
	}

	// handle events
	switch a.Mode {
	case inGame: // el juego se encuentra corriendo
		// si el juego termina reiniciar los valores y mostrar el menu principal
		if a.Game.Done {
			a.Game.Done = false
			a.Mode = menu
			return a, nil
		}
		// actualizar el juego
		a.Game, cmd = a.Game.Update(msg)
		return a, cmd

	case menu: // actualizar el menu
		a.Menu, cmd = a.Menu.Update(msg)
		// si desde el menu se selecciono algo
		if a.Menu.Selected {
			return a.selectMode()
		}
		return a, cmd

	case local: // actualizar el menu
		a.LocalQuotes, cmd = a.LocalQuotes.Update(msg)
		// si desde el menu se selecciono algo
		if a.LocalQuotes.Selected {
			return a.selectLocalQuote()
		}
		return a, cmd
	}
	return a, cmd
}

// selecciona la vista dependiendo del estado de la aplicacion
func (m App) View() string {
	switch m.Mode {
	case inGame: // mostrar juego
		return docStyle.Render(m.Game.View())

	case resumen: // mostrar juego
		return docStyle.Render(m.Game.View())

	default: // mostrar menu
		return docStyle.Render(m.Menu.View())
	}
}

/*
	triggered when and option is selected in the main menu.

Handles the App state and sets the correct mode
*/
func (a App) selectMode() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	a.Menu.Selected = false
	// change game mode
	switch a.Menu.List.SelectedItem().FilterValue() {
	case "jugar": // crear una nueva instancia del typer
		a.Mode = inGame
		a.Game = NewTyper(a.appWith)
		cmd = a.Game.Init()
		return a, cmd

	case "offline": // nuevo typer pero con el archivo local
		a.Mode = inGame
		a.Game = NewTyper(a.appWith)
		cmd = a.Game.Init()
		return a, cmd

	case "cargar":
		return a, tea.Quit
	}
	return a, cmd
}

func (a App) selectLocalQuote() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	a.LocalQuotes.Selected = false
	a.Game = NewTyper(a.appWith, a.LocalQuotes.Selection)
	return a, cmd
}
