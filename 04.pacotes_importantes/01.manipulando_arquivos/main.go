package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Olá, mundo!")
	tamanho, err := f.Write([]byte("Olá, mundo!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	f.Close()

	// Lendo arquivo
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Lendo arquivo...")
	fmt.Println(string(file))

	// Ler arquivo aos poucos - Buffer
	file2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
