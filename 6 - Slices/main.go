package main

import "fmt"

// Slices

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// len = 10, cap = 10, [10 20 30 40 50 60 70 80 90 100]
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)

	// :0 apaga os elementos do slice a direita do indice 0
	fmt.Printf("len = %d, cap = %d, %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// :4 apaga os elementos do slice a direita do indice 4
	fmt.Printf("len = %d, cap = %d, %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// 2: apaga os elementos do slice a esquerda do indice 2
	fmt.Printf("len = %d, cap = %d, %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// inserindo um elemento no slice, o cap do slice dobra de tamanho por causa do array interno, logo o slice aponta para um novo array com o dobro de tamanho
	s = append(s, 110)
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)
}
