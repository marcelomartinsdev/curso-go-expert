package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCEP)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err)
	}
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Server!"))
}
