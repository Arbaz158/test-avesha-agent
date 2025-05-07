package main

import (
	"fmt"
	"test-avesha-agent/math"
)

func main() {
	// Example usage of math operations
	a, b := 10, 5

	sum := math.Add(a, b)
	difference := math.Subtract(a, b)

	fmt.Printf("Addition: %d + %d = %d\n", a, b, sum)
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, difference)
}
