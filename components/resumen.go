package components

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
