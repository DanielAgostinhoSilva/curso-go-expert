package main

const a = "Hello, World!"

var (
	b bool
	c int
	d string
	e float64
)

func main() {
	a := "X"

	println(a)

	println(b)
	b = true
	println(b)

	println(c)
	c = 10
	println(c)

	println(d)
	d = "Daniel"
	println(d)

	println(e)
	e = 10.5
	println(e)
}
