package main

import "time"

func worker(workerId int, data <-chan int) {
	for x := range data {
		println("Worker", workerId, "received", x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtdWorkers := 1000000

	// inicializa os workers
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10000000; i++ {
		data <- i
	}
}
