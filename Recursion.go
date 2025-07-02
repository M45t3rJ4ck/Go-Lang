package main

import (
	"fmt"
)

// Recursion - the act in which the function calls itself (Must be a condition inside the function that ends the function calling itself.).
func factorial(num uint64) uint64 {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	pl := fmt.Println

	pl("Factorial of 16:", factorial(16))
}
