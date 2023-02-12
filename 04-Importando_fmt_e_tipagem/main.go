package main

import "fmt"

const a = "Hello, World!"

type ID int

// global scope variable declaration
var (
	// declaring and assigning a variable
	b bool = true
	c int
	d string
	e float64
	f ID = 1
)

func main() {
	fmt.Printf("O tipo de A é %T ", a)
	fmt.Printf("e o valor é %v\n", a)

	fmt.Printf("O tipo de B é %T ", b)
	fmt.Printf("e o valor é %v\n", b)

	fmt.Printf("O tipo de C é %T ", c)
	fmt.Printf("e o valor é %v\n", c)

	fmt.Printf("O tipo de D é %T ", d)
	fmt.Printf("e o valor é %v\n", d)

	fmt.Printf("O tipo de E é %T ", e)
	fmt.Printf("e o valor é %v\n", e)

	fmt.Printf("O tipo de F é %T ", f)
	fmt.Printf("e o valor é %v\n", f)
}
