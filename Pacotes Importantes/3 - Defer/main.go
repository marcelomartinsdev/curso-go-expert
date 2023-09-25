package main

import (
	"fmt"
	"io"
	"net/http"
)

// O defer executa a funcao apos o retorno da funcao atual, basicamente ele adia a execucao da funcao
func main() {
	req, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}
	// O defer sempre sera chamado apenas por ultimo, ou seja, nesse caso, abrimos um arquivo e para nao deixarmos de fechar, usamos o defer de cara para garantir que apos tudo ser feito, o arquivo sera fechado
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

	// Exemplo simples - Output será Segunda - Terceira - Primeira
	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
