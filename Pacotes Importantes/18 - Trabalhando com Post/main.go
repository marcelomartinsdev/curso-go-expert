package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

// Tempo limite para requisição HTTP (em segundos) ser concluída, caso contrario será cancelada com o erro "net/http: request canceled (Client.Timeout exceeded while awaiting headers)"

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "marcelo"`))
	resp, err := c.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
