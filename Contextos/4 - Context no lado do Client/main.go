package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// Crie um contexto com um prazo de 3 segundos.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel() // Certifique-se de cancelar o contexto quando a função main() terminar.

	// Crie uma nova solicitação HTTP com o contexto criado.
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	// Envie a solicitação HTTP e obtenha a resposta.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // Certifique-se de fechar o corpo da resposta quando terminar.

	// Copie o corpo da resposta para a saída padrão (stdout), neste caso, o console.
	io.Copy(os.Stdout, res.Body)
}
