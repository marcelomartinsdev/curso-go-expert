package main

func main() {
	// Memoria -> Endereco -> Valor

	// Ponteiros são tipos de dados que armazenam o endereço de memória de outro valor.

	// variavel -> ponteiro que tem um endereco de memoria -> valor
	a := 10
	// Para criar um ponteiro, basta utilizar o operador & antes do nome da variável.
	var ponteiro *int = &a
	// Para acessar o valor de um ponteiro, utilizamos o operador * antes do nome da variável.
	println(ponteiro)
	*ponteiro = 20
	println(a)
	b := &a
	println(*b)
	*b = 30
	println(a)
	c := &a
	*c = 40
	println(*c)
	println(a)
}
