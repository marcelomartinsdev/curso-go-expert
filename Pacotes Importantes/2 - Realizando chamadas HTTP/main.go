package main

import (
	"io"
	"net/http"
)

// para fazer chamadas http no go, basta usar o pacote net/http e a funcao Get do pacote http

func main() {
	req, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close()
}
