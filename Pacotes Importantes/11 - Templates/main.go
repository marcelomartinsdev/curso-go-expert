package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

// O template em Go é uma string que pode ser pré-compilada e executada várias vezes, o que é muito útil para gerar documentos HTML, por exemplo. O template é uma string que pode conter variáveis, que são representadas por {{.NomeDaVariavel}}. O template é compilado e executado com o método Execute, que recebe um io.Writer e um objeto que será usado para substituir as variáveis do template.
func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
