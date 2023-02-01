package components

import tea "github.com/charmbracelet/bubbletea"

type ResumenForm struct {
	assserts int
	errors   int
}

type Resumen struct {
	info ResumenForm
}

func NewResumen() Resumen {
	return Resumen{
		info: ResumenForm{
			assserts: 1,
			errors:   1,
		},
	}
}

// despliega la vista del resumen
func (r Resumen) View() string {
	return ""
}

func (r Resumen) Update() (Resumen, tea.Cmd) {
	return Resumen{}, nil
}
