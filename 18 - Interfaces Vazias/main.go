package main

import "fmt"

// Interfaces vazias são interfaces que não possuem nenhum método. Elas são usadas para representar qualquer tipo de dado, pois qualquer tipo de dado implementa uma interface vazia.

func main() {

	var x interface{} = 10
	var y interface{} = "Olá"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor e %v\n", t, t)
}
