package main

import "fmt"

// Structs são tipos de dados que permitem agrupar valores de tipos diferentes.

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Interfaces são tipos de dados que definem um conjunto de métodos que uma struct deve implementar para que possa ser considerada do tipo da interface.

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	fmt.Printf("A empresa %s agora está desativada.\n", e.Nome)
}

// Para que  todos Clientes tenham um endereço, podemos criar uma struct Cliente que contém um campo do tipo Endereco.
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

// Para criar métodos em structs, basta criar uma função com um receiver do tipo da struct.
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s agora está desativado.\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
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
	Desativacao(marcelo)
	minhaEmpresa := Empresa{
		Nome: "Caelum",
	}
	Desativacao(minhaEmpresa)

	fmt.Printf("O nome do cliente é %s e ele tem %d anos. Ativo: %t\n", marcelo.Nome, marcelo.Idade, marcelo.Ativo)
}
