package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/blog", blog{title: "Blog de Marcelo"})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Server!"))
	})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Server 2!"))
	})
	http.ListenAndServe(":8081", mux2)
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("b.title!"))
}
