package tools

import (
	"io/ioutil"
	"os"
)

// Calculate a Fibonacci number
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}