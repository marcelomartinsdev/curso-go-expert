package main

import (
	"os"
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
	t := template.Must(template.New("template.html").ParseFiles("Pacotes Importantes/13 - Templates com Arquivo Externo/template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 80},
		{"Python", 120},
		{"JavaScript", 180},
	})
	if err != nil {
		panic(err)
	}
}
