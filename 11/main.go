package main

import "fmt"

type Endereco struct {
	Logradouro string
	Cidade     string
	Estado     string
	Numero     int
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // Compondo
	//Address  Endereco // Atributo do tipo Endereco
}

func main() {

	williams := Cliente{
		Nome:  "Williams",
		Idade: 29,
		Ativo: true,
	}

	williams.Cidade = "Camaragibe"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", williams.Nome, williams.Idade, williams.Ativo)
	fmt.Println(williams)
}
