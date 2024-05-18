package main

import (
	"fmt"
	"os"

	"calculator/ops/add"
	"calculator/ops/subtract"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Overload()
	fmt.Println(os.Getenv("SYS_NAME"))
	fmt.Println("Add 10 and 10", add.AddTwoNumbers(10, 10))
	fmt.Println("Subtract 10 and 10", subtract.SubtractTwoNumbers(10, 10))
}
