package main

import (
	"fmt"

	"github.com/learngo/dockeredgo/src/server"
)

func main() {
	fmt.Println("Running Main")
	defer fmt.Println("Exiting Main")

	server.RunServer()
}
