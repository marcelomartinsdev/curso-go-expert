package main

import (
	"io"
	"net/http"
	"time"
)

// Tempo limite para requisição HTTP (em segundos) ser concluída, caso contrario será cancelada com o erro "net/http: request canceled (Client.Timeout exceeded while awaiting headers)"

func main() {
	c := http.Client{Timeout: time.Microsecond}
	resp, err := c.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
