package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas.")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
