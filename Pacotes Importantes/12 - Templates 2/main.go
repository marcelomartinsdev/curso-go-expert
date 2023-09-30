package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

// O template.must é uma função que retorna um template e um erro, mas como o erro é sempre nil, podemos ignorá-lo, se o erro não for nil, o programa para.
// Com o template.must criamos o template usando o template.new, que recebe o nome do template e o template em si, que é uma string. O template é uma string que pode conter variáveis, que são representadas por {{.NomeDaVariavel}}. Logo apos usamos o template.parse, que recebe o template e retorna um template e um erro, mas como o erro é sempre nil, podemos ignorá-lo.
// Com o template pronto, podemos executá-lo, passando o os.Stdout como saída e o curso como dados.

func main() {
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
