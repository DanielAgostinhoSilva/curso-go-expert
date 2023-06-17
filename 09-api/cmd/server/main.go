package main

import "github.com/DanielAgostinhoSilva/curso-go-expert/api/77-criando-arquivos-de-configuracao/configs"

func main() {
	config, _ := configs.LoadConfig("./cmd/server/.env")
	println(config.DBDriver)
}
