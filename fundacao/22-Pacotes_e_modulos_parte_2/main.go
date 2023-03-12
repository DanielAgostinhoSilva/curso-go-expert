package main

import (
	"curso-go-expert/22-Pacotes_e_modulos_parte_2/matematica"
	"fmt"
)

func main() {
	resultadoSoma := matematica.Soma(10, 20)

	fmt.Printf("Resultado: %v\n", resultadoSoma)
}
