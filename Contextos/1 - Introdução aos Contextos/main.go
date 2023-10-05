package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Crie um contexto com cancelamento para controlar o servidor HTTP.
	// O contexto permite controlar a execução e o encerramento de tarefas assíncronas.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Certifique-se de cancelar o contexto quando o programa terminar.

	// Inicie o servidor HTTP em uma goroutine.
	go startServer(ctx)

	// Espere um pouco antes de cancelar o contexto e desligar o servidor.
	time.Sleep(5 * time.Second)
	fmt.Println("Desligando o servidor HTTP...")
	cancel() // Isso sinaliza para o servidor que ele deve parar.

	// Aguarde até que o servidor tenha sido interrompido completamente.
	time.Sleep(2 * time.Second)
	fmt.Println("O programa foi encerrado.")
}

func startServer(ctx context.Context) {
	// Crie um multiplexador de roteamento para lidar com solicitações HTTP.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá, Mundo!")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Use o contexto para controlar a execução do servidor.
	go func() {
		fmt.Println("Iniciando o servidor HTTP...")
		// Inicie o servidor HTTP em segundo plano e gerencie erros.
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Erro no servidor HTTP: %v\n", err)
		}
	}()

	// Aguarde o cancelamento do contexto para desligar o servidor.
	<-ctx.Done() // Isso bloqueará até que o contexto seja cancelado.

	// Desligue o servidor HTTP com um prazo para encerrar as conexões existentes.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Desligando o servidor HTTP...")
	// Tente desligar o servidor HTTP de forma limpa, dando um tempo limite.
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Erro ao desligar o servidor HTTP: %v\n", err)
	}
}
