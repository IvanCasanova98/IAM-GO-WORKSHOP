package main

import (
	"errors"
	"fmt"
)

var DivisionByZero = errors.New("Division por cero")

func safedivide(a int, b int) (int, error) {
	if b == 0 {
		return 0, DivisionByZero
	}
	return a / b, nil
}

func main() {
	a, b := 1, 5
	c, err := safedivide(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d entre %d es %d \n", a, b, c)
	a, b = 10, 0
	c, err = safedivide(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d entre %d es %d \n", a, b, c)
}


func error_handler() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    panic_attack(0)
    fmt.Println("Returned normally from g.")
}

func panic_attack(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    panic_attack(i + 1)
}