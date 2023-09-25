package main

import "fmt"

// Closures tambem sao conhecidas como funcoes anonimas, sao funcoes que nao possuem nome e podem ser atribuidas a uma variavel

func main() {
	total := func() int {
		return sum(5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
