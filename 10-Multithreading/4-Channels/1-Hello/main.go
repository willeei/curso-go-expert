package main

// Thread 1
func main() {
	canal := make(chan string) // Canal vazio

	// Thread 2
	go func() {
		canal <- "Hello" // Canal cheio
	}()

	// Thread 1
	msg := <-canal // Canal esvaziado
	println(msg)
}
