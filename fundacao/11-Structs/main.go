package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	daniel := Cliente{
		Nome:  "Daniel",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t \n", daniel.Nome, daniel.Idade, daniel.Ativo)
	fmt.Println("Struct: ", daniel)

	daniel.Ativo = false
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t \n", daniel.Nome, daniel.Idade, daniel.Ativo)
}
