package main

import (
	"fmt"
)

// funcoes variadicas sao funcoes que recebem um numero variavel de argumentos

func main() {
	fmt.Println(sum(5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5))
}

// a funcao sum recebe um numero variavel de argumentos do tipo int, percorre o slice de inteiros e retorna a soma de todos os elementos
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
