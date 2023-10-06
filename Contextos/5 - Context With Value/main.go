package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}

// context.WithValue é uma função em Go que permite associar valores arbitrários a um contexto.
//Esses valores podem ser recuperados em qualquer ponto do programa que tenha acesso ao contexto.
//É especialmente útil para passar informações específicas da aplicação, como autenticação, identificadores de sessão ou configurações,
//através de várias chamadas de função sem a necessidade de passá-los explicitamente como argumentos.
