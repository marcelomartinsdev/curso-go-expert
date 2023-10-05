package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()                            // Contexto pai vazio
	ctx, cancel := context.WithTimeout(ctx, time.Second*3) // Contexto filho com timeout de 3 segundos
	defer cancel()                                         // Cancela o contexto filho quando a função main terminar
	bookHotel(ctx)                                         // Passa o contexto filho para a função bookHotel e executa
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): // Se o contexto filho for cancelado, o Done() será executado e o case será executado
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(time.Second * 5): // Se o contexto filho não for cancelado, o After() será executado e o case será executado após 5 segundos
		fmt.Println("Hotel booked.")
	}
}
