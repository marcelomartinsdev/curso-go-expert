package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("Pacotes Importantes/10 - File Server/public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Blog!"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
