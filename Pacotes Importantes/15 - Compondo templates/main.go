package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

// Nao faz muito sentido passar todos os valores para o parse do template
// Faz mais sentido passar o arquivo externo

func main() {

	templates := []string{
		"Pacotes Importantes/15 - Compondo templates/header.html",
		"Pacotes Importantes/15 - Compondo templates/content.html",
		"Pacotes Importantes/15 - Compondo templates/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"Go", 10},
			{"Java", 10},
			{"Python", 10},
			{"JavaScript", 10},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)
}
