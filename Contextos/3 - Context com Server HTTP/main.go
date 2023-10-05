package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// Configura uma rota padrão para a raiz ("/") e um manipulador (handler) para ela.
	http.HandleFunc("/", handler)

	// Inicia um servidor HTTP que escuta na porta 8080 e utiliza o roteador (mux) padrão (nil).
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Obtém um contexto associado à requisição atual. O contexto permite o controle
	// de prazos e cancelamentos associados a esta requisição.
	ctx := r.Context()

	// Loga o início da requisição.
	log.Println("Request Iniciada!")

	// Atrasa a execução por 5 segundos (simulando uma operação demorada).
	select {
	case <-time.After(5 * time.Second):
		// Imprime no console (stdout) para fins de registro.
		log.Println("Request processada com sucesso!")

		// Escreve no navegador do cliente a resposta de sucesso.
		w.Write([]byte("Request processada com sucesso!"))

	case <-ctx.Done():
		// Se o contexto for cancelado antes que o atraso expire, isso pode ocorrer
		// se o cliente cancelar a solicitação, por exemplo, devido a um tempo limite,
		// então tratamos isso aqui.

		// Imprime no console (stdout) para fins de registro.
		log.Println("Request cancelada pelo cliente!")

		// Escreve no navegador do cliente uma resposta de erro indicando que a
		// solicitação foi cancelada pelo cliente, usando o código de status HTTP 408
		// (Request Timeout).
		http.Error(w, "Request cancelada pelo cliente!", http.StatusRequestTimeout)
	}

	// Loga o final da requisição (isso é executado independentemente de sucesso
	// ou cancelamento).
	defer log.Println("Request Finalizada!")
}
