package main

import "fmt"

func init() {
	fmt.Println("Se ejecuta init")
}

func main() {
	var a int
	fmt.Println("a: ", a)

	var b = 10
	fmt.Println("b: ", b)

	c := 10
	fmt.Println("c: ", c)

	// Tipos Basicos
	var _ bool
	var _ int
	var _ string
	var _ float32
	var _ complex128

	// Tipos compuestos
	var array [4]int = [4]int{1, 2, 3, 4}

	var slice []int = []int{1, 2, 3, 4, 5, 6, 7}

	slice = append(slice, 8)

	var mapeo map[string]string = map[string]string{"name": "ivan"}

	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
	fmt.Println("mapeo: ", mapeo)

	// Punteros
	var puntero = new(int)

	fmt.Println("direccion de la variable a la que apunto: ", puntero)
	fmt.Println("direccion del puntero: ", &puntero)
	fmt.Println("valor?: ", *puntero)

	var valor int = 22

	fmt.Println("direccion del valor: ", &valor)

	puntero = &valor

	fmt.Println("valor: ", *puntero)

	duplicar(puntero)

	fmt.Println("duplicado: ", puntero)

}

func duplicar(n *int) {
	*n = *n * 2
}
