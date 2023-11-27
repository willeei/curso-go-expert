package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	williams := Cliente{
		Nome:  "Williams",
		Idade: 29,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", williams.Nome, williams.Idade, williams.Ativo)
}
