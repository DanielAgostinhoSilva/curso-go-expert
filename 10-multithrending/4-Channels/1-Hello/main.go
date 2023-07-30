package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // canal vazio

	// Thread 2
	go func() {
		channel <- "Hello, World!" // canal cheio
	}()

	// Thread 1
	msg := <-channel // esvazia o canal
	fmt.Println(msg)
}
