package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

const url = "https://api.quotable.io/random?minLength=120"

type Cuote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
	Length  int    `json:"length"`
	Splited []string
}

// Retorna una nueva cita aleatoria de internet utilizando quotable
func NuevaCita() (*Cuote, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body Cuote
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	body.Splited = strings.Split(body.Content, " ")
	return &body, nil
}
