package main

import (
	"24-Instalando_pacotes/matematica"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	resultadoSoma := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Printf("Resultado: %v\n", resultadoSoma)

	fmt.Println(matematica.A)
	fmt.Println(carro.Marca)

	carro.Andar()

	fmt.Println(uuid.New())
}
