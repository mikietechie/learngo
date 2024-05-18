package main

import "fmt"

func CalculateInterest(p, r, t float64) float64 {
	return p * r * t
}

func main() {
	// 3. Write a C program to calculate Compound Interest.
	fmt.Println("Interest when Principal=100, rate=10%, time=3 is\t:\n", CalculateInterest(100, 0.10, 3))
}
