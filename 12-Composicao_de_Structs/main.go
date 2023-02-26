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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s for desativado", c.Nome)
}

func main() {
	daniel := Cliente{
		Nome:  "Daniel",
		Idade: 30,
		Ativo: true,
	}

	daniel.Desativar()
}
