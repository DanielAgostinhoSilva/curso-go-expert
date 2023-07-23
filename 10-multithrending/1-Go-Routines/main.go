package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%2.d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("A")
	go task("B")
	time.Sleep(15 * time.Second)
}
