package main

import (
	"request"
	"fmt"
	"log"
)

func main() {
	pelicula := Pelicula{"Spiderman", 140, "Accion"}

	payload, err := json.Marshal(&emp1)
	if err != nil {
		log.Fatal("No se pudo serializar")
	}
	fmt.Println(string(payload))

	tony := `{"tittle":"Superman","duration": 100 ,"genre":"Comedy"}`
	var pelicula_raw Pelicula
	err = json.Unmarshal([]byte(tony), &emp3)
	if err != nil {
		log.Fatal("No se pudo deserializar")
	}
	fmt.Println(emp3)
}