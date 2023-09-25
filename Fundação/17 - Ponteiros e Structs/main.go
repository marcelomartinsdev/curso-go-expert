package main

type Conta struct {
	saldo int
}

/* NewConta é um construtor de Conta, porque retorna um ponteiro para Conta, e constrói uma Conta com saldo 0

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

*/

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{
		saldo: 100,
	}
	conta.simular(200)
	println(conta.saldo)
}
