package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer")
	defer fmt.Println("Defer é uma instrução que adia a execução de uma função até o final do escopo atual.")
	fmt.Println("Defer é comumente usado para garantir que um recurso seja fechado antes de sair de uma função.")
}
