package main

import (
	"fmt"
	"github.com/DanielAgostinhoSilva/curso-go-expert/packaging/68-exportacaoes-de-pacotes/math"
)

func main() {
	a := math.NewMath(1, 2)
	fmt.Println(a.Add())
}
