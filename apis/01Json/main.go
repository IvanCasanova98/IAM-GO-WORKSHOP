package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pelicula struct {
	Titulo   string `json:"tittle,omitempty"`
	Duracion int    `json:"duration,omitempty"`
	Genero   string `json:"genre,omitempty"`
}

func main() {
	pelicula := Pelicula{"Spiderman", 140, "Accion"}

	payload, err := json.Marshal(&pelicula)
	if err != nil {
		log.Fatal("No se pudo serializar")
	}
	fmt.Println(string(payload))

	tony := `{"tittle":"Superman","duration": 100 ,"genre":"Comedy"}`
	var peliculaRaw Pelicula
	err = json.Unmarshal([]byte(tony), &peliculaRaw)
	if err != nil {
		log.Fatal("No se pudo deserializar")
	}
	fmt.Println(peliculaRaw)
}