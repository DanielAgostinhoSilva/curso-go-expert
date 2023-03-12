package main

const a = "Hello, World!"

// global scope variable declaration
var (
	// declaring and assigning a variable
	b bool = true
	c int
	d string
	e float64
)

func main() {
	// shorthand declaration
	f := "daniel"
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
