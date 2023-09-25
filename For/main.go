package main

func main() {

	// For padrao
	for i := 0; i < 10; i++ {
		println(i)
	}

	// Percorrer Slices e Vetores
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}

	// Parecido com o While
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	// Loop Infinito
	for {
		println("Hello, World!")
	}
}
