package main

import "fmt"

// receive-only channel
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// send-only channel
func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
