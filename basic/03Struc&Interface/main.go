package main

import "fmt"

type Entretenimiento interface {
	Titulo() string
	Duracion() int
}

func main() {

	var supermanMovie = &Pelicula{titulo: "Superman", duracion: 150, genero: "accion"}
	fmt.Println(supermanMovie.titulo)

	var emptyMovie = &Pelicula{}
	fmt.Println(emptyMovie.titulo)

	(*emptyMovie).titulo = (*supermanMovie).titulo

	var serieVacia = &Serie{}

	var cineVacio = &Cine{}

	var Backlog = []Entretenimiento{supermanMovie, emptyMovie, serieVacia, cineVacio}

	for _, v := range Backlog {
		fmt.Println(v.Titulo())
	}

}

type Pelicula struct {
	titulo   string
	duracion int
	genero   string
}

func (peli *Pelicula) Titulo() string {
	return fmt.Sprintf("Soy la pelicula: %s", peli.titulo)
}

func (peli *Pelicula) Duracion() int {
	return peli.duracion
}

type Serie struct {
	nombre    string
	duracion  int
	capitulos int
	genero    string
}

func (serie *Serie) Titulo() string {
	return fmt.Sprintf("Soy la pelicula: %s", serie.nombre)
}

func (serie *Serie) Duracion() int {
	return serie.capitulos * serie.duracion
}




type Cine struct{
	Pelicula
	nombre string
}