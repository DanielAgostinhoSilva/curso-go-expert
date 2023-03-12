package main

import (
	"curso-go-expert/23-Pacotes_e_modulos_parte_3/matematica"
	"fmt"
)

func main() {
	resultadoSoma := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Printf("Resultado: %v\n", resultadoSoma)

	fmt.Println(matematica.A)
	fmt.Println(carro.Marca)

	carro.Andar()
}
