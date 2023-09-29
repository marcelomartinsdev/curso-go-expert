package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err)
	}
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Contant-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Server!"))
}

// Caso a URL da requisição for diferente de "/", ira retornar status 404 Not Found
// Se a Query passada em ?cep for vazia ira retornar status 400 Bad Request
// Caso contrario, ira retornar status 200 OK e ira escrever Hello Server! na resposta da requisicao
