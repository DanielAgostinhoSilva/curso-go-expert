package main

import (
	"fmt"
	"github.com/DanielAgostinhoSilva/curso-go-expert/packaging/67-acessando-pacotes-criados/math"
)

func main() {
	a := math.Math{A: 1, B: 2}
	fmt.Println(a.Add())
}
