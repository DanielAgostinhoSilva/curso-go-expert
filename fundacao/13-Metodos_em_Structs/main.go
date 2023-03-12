package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
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

	daniel.Cidade = "Sao Paulo"
	fmt.Println("Struct: ", daniel)

	daniel.Endereco.Cidade = "Sao Paulo 2"
	fmt.Println("Struct: ", daniel)
}
