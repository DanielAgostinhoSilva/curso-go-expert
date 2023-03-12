package main

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
	// shorthand declaration
	a := "daniel"
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
