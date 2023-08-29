package main

import "fmt"

// Structs são tipos de dados que permitem agrupar valores de tipos diferentes.

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Para que  todos Clientes tenham um endereço, podemos criar uma struct Cliente que contém um campo do tipo Endereco.
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

// Para acessar os valores de uma struct, utilizamos o operador ponto (.) após o nome da variável.

func main() {

	marcelo := Cliente{
		Nome:  "Marcelo",
		Idade: 30,
		Ativo: true,
	}

	marcelo.Estado = "SP"
	marcelo.Numero = 123
	marcelo.Ativo = false

	fmt.Printf("O nome do cliente é %s e ele tem %d anos. Ativo: %t\n", marcelo.Nome, marcelo.Idade, marcelo.Ativo)
}
