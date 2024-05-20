package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		Curso{"Go", 40},
		Curso{"Python", 45},
		Curso{"Java", 60},
	})
	if err != nil {
		panic(err)
	}
}
