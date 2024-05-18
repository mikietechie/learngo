package main

import (
	"fmt"

	"github.com/learngo/rediscachehttpserver/src/server"
)

func main() {
	fmt.Println("Running Main")
	defer fmt.Println("Exiting Main")

	server.RunServer()
}
