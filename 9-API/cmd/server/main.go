package main

import "github.com/willbrr.dev/goexpert/9-API/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
