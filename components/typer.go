package components

import (
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elias-gill/type_game/utils"
)

// modelo basico del typer
type Typer struct {
	cita          *utils.Quote
	timer         stopwatch.Model
	textArea      textarea.Model
	coloredOutput []string
	userInputs    []string
	pos           int // posicion de la palabra en la cita
	outputSize    int
	Done          bool
}

/*
Retorna una nueva instancia del Typer (juego de escribir).

El "fraceId" corresponde al identificador de la cita local ("id"). Si el valor proporcionado es nulo, entonces se retorna un typer con
una cita sacada de internet.
*/
func NewTyper(width int, quote ...utils.Quote) Typer {
	done := false
	var cita *utils.Quote
	if quote == nil {
		var err error
		cita, err = utils.NuevaCitaOnline()
		if err != nil {
			done = true
		}
	}

	// crear un array auxiliar con espacios en blanco
	aux := []string{}
	for i := 0; i < len(cita.Splited); i++ {
		aux = append(aux, "")
	}

	return Typer{
		Done:          done,
		textArea:      newTextArea(),
		timer:         stopwatch.NewWithInterval(time.Second),
		outputSize:    width, // window width
		cita:          cita,
		coloredOutput: cita.Splited,
		userInputs:    aux,
	}
}

func (t Typer) Init() tea.Cmd {
	return t.timer.Init()
}

func (t Typer) View() string {
	s := "\n\n\t\t"
	// fomatear cita
	for i, v := range t.coloredOutput {
		/* Cada palabra contiene en promedio 10 letras (ingles) entonces se calcula la cantitdad de
		   palabras que entra en un outputSize y cada multiplo se agrega un salto de linea */
		size := int(t.outputSize / 8)
		if size == 0 {
			size = 1
		}
		if i%size == 0 {
			s += "\n\t\t"
		}
		s += v + " "
	}
	s += authorStyle.Render("\n\t\t- '" + t.cita.Author + "'")
	s += "\n\n" + t.textArea.View()
	s += "\n\t\t\t" + t.timer.View()
	return s
}

// Se encarga de actualizar el texto en pantalla y de colorear conforme el usuario escribe
func (t Typer) Update(msg tea.Msg) (Typer, tea.Cmd) {
	var cmd tea.Cmd
	t.timer, cmd = t.timer.Update(msg)

	switch msg := msg.(type) {
	// handle when the window is resized
	case tea.WindowSizeMsg:
		h, _ := docStyle.GetFrameSize()
		t.outputSize = msg.Width - h
		return t, cmd

	case tea.KeyMsg:
		// teclas especiales
		switch msg.String() {
		case "alt+backspace", "backspace", "ctrl+w": // volver una palabra atras
			if len(t.textArea.Value()) == 0 && t.pos > 0 {
				t.coloredOutput[t.pos] = t.cita.Splited[t.pos]
				t.pos--
				t.textArea.SetValue(t.userInputs[t.pos])
				t.colorearPalActual()
				return t, cmd
			}

		case "ctrl+c": // salir del programa
			return t, tea.Quit

		case "esc": // salir al menu
			t.Done = true
			t.timer.Stop()
			return t, tea.ClearScreen

		case " ": // colorear y pasar a la siguiente palabra
			return t.colorearOutput(), cmd
		}
	}

	// actualizar el textArea con un input normal
	t.textArea, _ = t.textArea.Update(msg)
	// colorear las letras de la palabra actual
	t.colorearPalActual()
	return t, cmd
}

/*
	Colorea el input de las palabras ya terminadas (una palabra se considera terminada cuando

se preciona la tecla espacio)
*/
func (t Typer) colorearOutput() Typer {
	t.userInputs[t.pos] = t.textArea.Value()
	// terminar el juego cuando se llega a la ultima palabra
	if t.pos == len(t.cita.Splited)-1 {
		t.Done = true
		t.timer.Stop()
		return t
	}

	// pintar las palabras que estan bien y las que estan mal
	if t.cita.Splited[t.pos] == t.textArea.Value() {
		t.coloredOutput[t.pos] = goodStyle.Render(t.cita.Splited[t.pos])
	} else {
		t.coloredOutput[t.pos] = badStyle.Render(t.cita.Splited[t.pos])
	}
	t.textArea.Reset()
	t.pos += 1
	return t
}

/*
	colorea la palabra que se esta escribiendo actualmente letra por letra dependiendo de

lo que el usuario escribe
*/
func (t Typer) colorearPalActual() {
	text := t.textArea.Value()
	s := ""
	palActual := t.cita.Splited[t.pos]

	if len(text) > len(palActual) {
		t.coloredOutput[t.pos] = badStyle.Render(t.cita.Splited[t.pos])
		return
	}

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
	// t.coloredOutput[t.pos] = s
}

// funcion para crear un nuevo text area (donde el usuario escribe)
func newTextArea() textarea.Model {
	// generar un nuevo TextArea
	ta := textarea.New()
	ta.Placeholder = ""
	ta.Focus()

	// estilo
	ta.Prompt = doneStyle.Render("\t\t??? ")
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle() // vaciar el estilo por defecto
	ta.ShowLineNumbers = false

	// presets generales
	ta.SetWidth(50)
	ta.SetHeight(3)
	ta.CharLimit = 280
	ta.KeyMap.InsertNewline.SetEnabled(false)
	return ta
}
