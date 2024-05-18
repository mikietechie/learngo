package main

import (
	"fmt"
	"math"
)

func IsPrime(num int) bool {
	limit := int(math.Sqrt(float64(num)) + 1)
	for i := 2; i < limit; i++ {
		if (num % i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	// 2. Write a Program to check whether a number is prime or not.
	for i := 0; i < 20; i++ {
		fmt.Println(i, IsPrime(i))
	}
}
