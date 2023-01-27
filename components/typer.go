package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	// "github.com/charmbracelet/bubbles/viewport"
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
	outputSize    int
	Done          bool
}

// retorna una nueva instancia del Typer (juego de escribir)
func NewTyper(width int, s ...string) Typer {
	done := false
	cita, err := utils.NuevaCita(s)
	if err != nil {
		done = true
	}

	t := Typer{
		Done:          done,
		textArea:      newTextArea(),
		outputSize:    width, // largo de la pantalla en caracteres
		cita:          cita,
		coloredOutput: strings.Split(cita.Content, " "),
	}
	return t
}

func (t Typer) Init() tea.Cmd {
	return nil
}

func (t Typer) View() string {
	s := "\n\n\t\t"
	// fomatear cita
	for i, v := range t.coloredOutput {
        /* Cada palabra contiene en promedio 10 letras (ingles) entonces se calcula la cantitdad de 
        palabras que entra en un outputSize y cada multiplo se agrega un salto de linea */
		if i%int(t.outputSize/8) == 0 {
			s += "\n\t\t"
		}
		s += v + " "
	}
	s += authorStyle.Render("\n\t\t- '" + t.cita.Author + "'")
	s += "\n\n"
	s += t.textArea.View()
	return s
}

// Se encarga de actualizar el texto en pantalla y de colorear conforme el usuario escribe
func (t Typer) Update(msg tea.Msg) (Typer, tea.Cmd) {
	switch msg := msg.(type) {

	// handle when the window is resized
	case tea.WindowSizeMsg:
		h, _ := docStyle.GetFrameSize()
		t.outputSize = msg.Width - h
		return t, nil

	case tea.KeyMsg:
		// teclas especiales
		switch msg.String() {
		case "ctrl+c": // salir del programa
			return t, tea.Quit

		case "esc": // salir al menu
			t.Done = true
			return t, tea.ClearScreen

		case " ": // colorear y pasar a la siguiente palabra
			return t.colorearOutput()
		}
	}

	// actualizar el textArea con un input normal
	var cmd tea.Cmd
	t.textArea, cmd = t.textArea.Update(msg)
	// colorear las letras de la palabra actual
	t.colorearPalActual()
	return t, cmd
}

/* Colorea el input de las palabras ya terminadas (una palabra se considera terminada cuando se preciona
la tecla espacio) */
func (t Typer) colorearOutput() (Typer, tea.Cmd) {
	// terminar el juego cuando se llega a la ultima palabra
	if t.pos == len(t.cita.Splited)-1 {
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

/* colorea la palabra que se esta escribiendo actualmente letra por letra dependiendo de lo que el usuario
escribe */
func (t Typer) colorearPalActual() {
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

func newTextArea() textarea.Model {
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
	return ta
}
