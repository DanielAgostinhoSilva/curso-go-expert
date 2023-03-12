package main

import "fmt"

func main() {
	salarios := map[string]int{"Daniel": 1000, "Joao": 2000, "Maria": 3000}
	delete(salarios, "Daniel")
	salarios["Dan"] = 5000

	//sal := make(map[string]int)
	//sal1 := map[string]int{}
	//sal["Daniel"] = 1000
	//sal1["Sandra"] = 2000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é de %d\n", salario)
	}

}
