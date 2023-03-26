package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)

	if err != nil {
		println(err)
	}
	println(string(res))

	conta2 := Conta{Numero: 2, Saldo: 100}

	err = json.NewEncoder(os.Stdout).Encode(conta2)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n":2, "s": 200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo)
}
