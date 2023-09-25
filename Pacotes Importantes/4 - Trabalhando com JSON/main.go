package main

import (
	"encoding/json"
	"os"
)

// Para usar o json, basta usar o Marshal, que transforma o objeto em um array de bytes e o Unmarshal, que transforma o array de bytes em um objeto

// Transformar uma struct em JSON
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 2000}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	// Para fazer o caso contrario, do JSON voltar para struct
	jsonPuro := []byte(`{"n": 2, "s": 200}`)
	var contaX Conta
	// Para alterar uma variavel é necessario ir no endereço que ela esta alocada e alterar, utilizando o &
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}
