package main

import (
	"fmt"
	"slices"
)

func main() {
	// 1. Find the largest number among the three numbers.
	threenumbers := []float64{1, 2, 3}
	fmt.Println(slices.Max(threenumbers))
}
