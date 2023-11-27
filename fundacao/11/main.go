package main

import "fmt"

type Endereco struct {
	Logradouro string
	Cidade     string
	Estado     string
	Numero     int
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // Compondo
	//Address  Endereco // Atributo do tipo Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
}

func main() {

	williams := Cliente{
		Nome:  "Williams",
		Idade: 29,
		Ativo: true,
	}

	williams.Cidade = "Camaragibe"
	williams.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", williams.Nome, williams.Idade, williams.Ativo)
	fmt.Println(williams)
}
