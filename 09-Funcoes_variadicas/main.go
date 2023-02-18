package main

import (
	"fmt"
)

func main() {
	fmt.Println(10, 20.30, 40, 50, 60, 70, 80, 90, 100)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
