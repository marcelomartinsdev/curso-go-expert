package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

// Nao faz muito sentido passar todos os valores para o parse do template
// Faz mais sentido passar o arquivo externo

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"Pacotes Importantes/16 - Mapeando funções nos Templates/header.html",
		"Pacotes Importantes/16 - Mapeando funções nos Templates/content.html",
		"Pacotes Importantes/16 - Mapeando funções nos Templates/footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 10},
		{"Java", 10},
		{"Python", 10},
		{"JavaScript", 10},
	})
	if err != nil {
		panic(err)
	}
}
