package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var soma, err = sum(10, 41)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println(soma)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
