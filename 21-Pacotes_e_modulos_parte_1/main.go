package main

import (
	"curso-go-expert/21-Pacotes_e_modulos_parte_1/matematica"
	"fmt"
)

func main() {
	resultadoSoma := matematica.Soma(10, 20)

	fmt.Printf("Resultado: %v\n", resultadoSoma)
}
