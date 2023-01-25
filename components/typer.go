package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elias-gill/type_game/utils"
)

// styles to colorate strings while the user is typing
var (
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

// modelo basico del typer
type Typer struct {
	cita          *utils.Cuote
	textArea      textarea.Model
	coloredOutput []string
	pos           int // posicion de la palabra en la cita
	asserts       int
	errors        int
	timer         Timer
	Done          bool
}

// retorna una nueva instancia del Typer (juego de escribir)
func NewTyper() *Typer {
	// generar un nuevo TextArea
	ta := textarea.New()
	ta.Placeholder = ""
	ta.Focus()

	// estilo
	ta.Prompt = doneStyle.Render("\t\t┃ ")
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle() // vaciar el estilo por defecto
	ta.ShowLineNumbers = false

	// presets generales
	ta.SetWidth(50)
	ta.SetHeight(3)
	ta.CharLimit = 280
	ta.KeyMap.InsertNewline.SetEnabled(false)

	// TODO: change behavior to not panic when the request is invalid
	cita, err := utils.NuevaCita()
	if err != nil {
		panic("bad request")
	}
	t := Typer{
		Done:          false,
		textArea:      ta,
		cita:          cita,
		coloredOutput: strings.Split(cita.Content, " "),
	}
	return &t
}

func (t Typer) Init() tea.Cmd {
	return t.timer.stopwatch.Start()
}

func (t Typer) View() string {
	var s = "\n\n\t\t"
	// fomatear cita
	for i, v := range t.coloredOutput {
		if i%20 == 0 {
			s += "\n\t\t"
		}
		s += v + " "
	}
	s += authorStyle.Render("\n\t\t- '" + t.cita.Author + "'")
	s += "\n\n"
	s += t.textArea.View()
	s += t.timer.View()
	return s
}

// Se encarga de actualizar el texto en pantalla y de colorear conforme el usuario escribe
func (t Typer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// ver si no se presiono ctrl+c
		case "ctrl+c":
			return t, tea.Quit

		case " ": // pasar a la siguiente palabra
			if t.pos == len(t.coloredOutput)-1 {
				t.Done = true
				return t, nil
			}
			// pintar las palabras que estan bien y las que estan mal
			if t.cita.Splited[t.pos] == t.textArea.Value() {
				t.coloredOutput[t.pos] = goodStyle.Render(t.cita.Splited[t.pos])
				t.asserts++
			} else {
				t.coloredOutput[t.pos] = badStyle.Render(t.cita.Splited[t.pos])
				t.errors++
			}
			t.textArea.Reset()
			t.pos++
			return t, nil
		}
	}

	// actualizar el textArea
	t.textArea, cmd = t.textArea.Update(msg)
	// colorear mientras se escribe
	t.colorearStrings()
	return t, cmd
}

/* colorea la palabra que se esta escribiendo actualmente letra por letra dependiendo de lo que el usuario
escribe */
func (t Typer) colorearStrings() {
	text := t.textArea.Value()
	s := ""
	palActual := t.cita.Splited[t.pos]

	// pintar las letras de la palabra actual
	for i := range palActual {
		// la letra actual
		if i == len(text) {
			s += doneStyle.Render(string(palActual[i]))
			continue
		}

		// evitar un overflow
		if i > len(text)-1 {
			s += string(palActual[i])
			continue
		}

		// pintar lo bien y lo mal
		if text[i] == palActual[i] {
			s += goodStyle.Render(string(rune(palActual[i])))
		} else {
			s += badStyle.Render(string(rune(palActual[i])))
		}
	}
	t.coloredOutput[t.pos] = s
}
