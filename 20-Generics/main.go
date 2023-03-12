package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Fulano 1": 1000, "Fulano 2": 2000, "Fulano 3": 3000}
	m2 := map[string]float64{"Fulano 1": 1000.10, "Fulano 2": 2000.20, "Fulano 3": 3000.30}
	m3 := map[string]MyNumber{"Fulano 1": 1000, "Fulano 2": 2000, "Fulano 3": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10))
}
