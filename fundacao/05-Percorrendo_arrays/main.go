package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor de do indice é %d e o valor é %d\n", i, v)
	}
}
