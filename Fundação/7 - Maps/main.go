package main

import "fmt"

// Maps (hash tables) são estruturas de dados que armazenam valores associados a chaves

func main() {
	salarios := map[string]int{"Marcelo": 1000, "João": 5000, "Maria": 2000}

	fmt.Println(salarios)
	fmt.Println(salarios["João"])
	delete(salarios, "João")
	salarios["João"] = 10000
	fmt.Println(salarios["João"])
	fmt.Println(salarios)

	// maps não são ordenados e não possuem tamanho fixo como arrays, o make é usado para criar um map vazio
	sal := make(map[string]int)
	sal["Marcelo"] = 999

	// abrindo e fechando chave para criar um map vazio
	sal1 := map[string]int{}
	sal1["Marcelo"] = 100000

	// para percorrer um map, usamos o for range

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	// caso nao queira usar o nome da chave, podemos usar o underline, que é um blank identifier (identificador) em go
	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
