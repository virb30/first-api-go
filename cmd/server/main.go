package main

import "github.com/virb30/first-api-go/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
