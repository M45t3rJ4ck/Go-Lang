package main

import "fmt"

// Functions don't have be associated with identifiers.

// Passing functions to functions.
func useFunc(f func(int, int) int, x, y int) {
	fmt.Println("Answer: ", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	intSum := func(x, y int) int { return x + y }

	fmt.Println("5 + 4 =", intSum(5, 4))

	// Closures can change functions outside of the function.
	samp1 := 1
	changeVar := func() { samp1 += 1 }
	changeVar()
	fmt.Println("Sample1:", samp1)

	// Passing functions to functions.
	useFunc(sumVals, 5, 8)
}
