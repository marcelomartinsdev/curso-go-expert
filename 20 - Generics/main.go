package main

// Generics em Go é um recurso que permite a criação de funções e estruturas de dados que podem ser usadas com diferentes tipos de dados.

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
	m := map[string]int{"Wesley": 1000, "Marcelo": 2000, "Joao": 3000}
	m2 := map[string]float64{"Wesley": 1000.31, "Marcelo": 2000.34, "Joao": 3000.31}
	m3 := map[string]MyNumber{"Wesley": 1000, "Marcelo": 2000, "Joao": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
