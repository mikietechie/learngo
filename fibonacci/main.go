package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Fibonacci")
	var a, b int64 = 0, 1
	//  No sufficient error handling here
	count, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < count; i++ {
		fmt.Println(a, b)
		a, b = b, a+b
	}
}
