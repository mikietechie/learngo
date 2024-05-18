package main

import (
	"fmt"

	"github.com/learngo/dockeredgo/src/functions"
)

func main() {
	fmt.Println("Running Main")
	defer fmt.Println("Exiting Main")
	functions.RunLogTimeForever()
}
