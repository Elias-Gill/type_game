package components

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func NewMainMenu() MainMenu {
	items := []list.Item{
		menuItem{Action: "jugar", Tit: "Jugar", Desc: "Citas famosas sacadas aleatoriamente de internet"},
		menuItem{Action: "offline", Tit: "Jugar offline", Desc: "Citas sacadas del banco de citas local"},
		menuItem{Action: "cargar", Tit: "Personalizado", Desc: "Escribe tu propia cita personalizada y guardala en el banco de citas local"},
	}

	m := MainMenu{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Menu Principal"
	return m
}

type menuItem struct {
	Tit, Desc, Action string
}

func (i menuItem) Title() string       { return i.Tit }
func (i menuItem) Description() string { return i.Desc }
func (i menuItem) FilterValue() string { return i.Action }

type MainMenu struct {
	List     list.Model
	Selected bool
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

// actualizar el modelo
func (m MainMenu) Update(msg tea.Msg) (MainMenu, tea.Cmd) {
	options := map[string]struct{}{"q": {}, "esc": {}, "ctrl+c": {}}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			m.Selected = true
		}

		// si la tecla precionada es una de las de salir
		_, keyExit := options[msg.String()]
		if keyExit {
			println("See you latter !!!")
			time.Sleep(time.Second * 1)
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd

}

// mostrar menu de seleccion
func (m MainMenu) View() string {
	return m.List.View()
}
