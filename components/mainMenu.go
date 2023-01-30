package components

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func NewMainMenu() mainMenu {
	items := []list.Item{
		Item{Action: "jugar", Tit: "Jugar", Desc: "Citas famosas sacadas aleatoriamente de internet"},
		Item{Action: "offline", Tit: "Jugar offline", Desc: "Citas sacadas del banco de citas local"},
		Item{Action: "cargar", Tit: "Personalizado", Desc: "Escribe tu propia cita personalizada y guardala en el banco de citas local"},
	}

	m := mainMenu{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Menu Principal"
	return m
}

type Item struct {
	Tit, Desc, Action string
}

func (i Item) Title() string       { return i.Tit }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Action }

type mainMenu struct {
	List list.Model
}

// INFO: no se utiliza
func (m mainMenu) Init() tea.Cmd {
	return nil
}

// actualizar el modelo
func (m mainMenu) Update(msg tea.Msg) (mainMenu, tea.Cmd) {
	options := map[string]struct{}{"q": {}, "esc": {}, "ctrl+c": {}}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		_, keyExit := options[msg.String()]
		if keyExit {
			println("See you latter !!!")
			time.Sleep(time.Second * 1)
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

// mostrar menu de seleccion
func (m mainMenu) View() string {
	return m.List.View()
}
