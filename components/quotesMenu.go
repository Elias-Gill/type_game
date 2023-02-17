package components

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/type_game/utils"
)

func NewQuotesMenu() QuotesMenu {
	items := []list.Item{
		quotesItem{quoteID: "jugar", tittle: "Jugar", author: "Citas famosas sacadas aleatoriamente de internet"},
		quotesItem{quoteID: "offline", tittle: "Jugar offline", author: "Citas sacadas del banco de citas local"},
		quotesItem{quoteID: "cargar", tittle: "Personalizado", author: "Escribe tu propia cita personalizada y guardala en el banco de citas local"},
	}

	m := QuotesMenu{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Menu Principal"
	return m
}

type quotesItem struct {
	tittle, author, quoteID string
}

func (i quotesItem) Title() string       { return i.tittle }
func (i quotesItem) Description() string { return i.author }
func (i quotesItem) FilterValue() string { return i.quoteID }

type QuotesMenu struct {
	List      list.Model
	Selected  bool
	Selection utils.Quote
	quoteList []utils.Quote
}

func (m QuotesMenu) Init() tea.Cmd {
	return nil
}

// actualizar el modelo
func (m QuotesMenu) Update(msg tea.Msg) (QuotesMenu, tea.Cmd) {
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
func (m QuotesMenu) View() string {
	return m.List.View()
}
