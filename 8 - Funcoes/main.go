package main

import (
	"errors"
	"fmt"
)

// Dados podem ser passados ou nao na entrada e saida de uma funcao

func main() {
	fmt.Println(sum(5, 5))

	valor, err := sum1(20, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}
}

// Caso as entradas possuam o mesmo tipo, podemos omitir o tipo da ultima variavel declarada
// Funcoes podem ter mais de um retorno e podem ser nomeadas ou nao
func sum(a, b int) (int, bool) {
	if a+b >= 30 {
		return a + b, true
	}
	return a + b, false
}

// Em go, podemos retornar um erro como segundo retorno de uma funcao e tratar esse erro no retorno da funcao
// errors.New cria um novo erro com a mensagem passada como parametro e retorna um ponteiro para esse erro (error) e nil para o segundo retorno (int) ja que nao queremos retornar nada.
func sum1(a, b int) (int, error) {
	if a+b >= 30 {
		return 0, errors.New("A soma é maior que 30")
	}
	return a + b, nil
}
