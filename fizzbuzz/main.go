package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		panic("No Args provided")
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("No Args")
	}
	for iterator := 0; iterator < count; iterator++ {
		fmt.Println(iterator)
		over3 := (iterator % 3) == 0
		over5 := (iterator % 5) == 0
		if over3 && over5 {
			fmt.Println("Fizz Buzz ", iterator)
		} else if over3 {
			fmt.Println("Fizz ", iterator)
		} else if over5 {
			fmt.Println("Buzz ", iterator)
		}
	}
}
