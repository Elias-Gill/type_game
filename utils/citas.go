package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

const url = "https://api.quotable.io/random?minLength=120"

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
	ID      string `json:"id"`
	Length  int    `json:"length"`
	Splited []string
}

type cuotes struct {
	Content []Quote `json:"cuotes"`
}

/* Dentro de S se envia el id de una cita contenidad localmente. Si el valor de s es nulo, entonces
se genera una nueva cita de internet utilizando quotable. */
func NuevaCitaOnline() (*Quote, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body Quote
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	body.Splited = strings.Split(body.Content, " ")
	return &body, nil
}

// WARNING: CAMBIAR EL SISTEMA DE ARCHIVOS
// Busca una cita local que contenga el ID proporcionado
func GetCitasLocales(id string) ([]*Quote, error) {
	dir := "~/palabras.json"
	file, err := os.Open(dir)
	if err != nil {
		panic("No se pudo encontrar el archivo de palabras locales " + dir)
	}

	var c cuotes
	json.NewDecoder(file).Decode(&c)

	for _, v := range c.Content {
		if v.ID == id {
            return []*Quote{}, nil
		}
	}
	panic("No se pudo encontrar la cita")
}
