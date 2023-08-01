package main

import "github.com/DanielAgostinhoSilva/curso-go-expert/11-eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	rabbitmq.Publish(ch, "Hello, World2!", "amq.direct")
}
