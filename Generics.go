package main

import (
	"fmt"
)

// More Constraints at pkg.go.dev/golang.org/x/exp/constraints
type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	pl := fmt.Println

	pl("5 + 4 =", getSumGen(5, 4))
	pl("5 + 4.7 =", getSumGen(5, 4.7))
	pl("5.6 + 4 =", getSumGen(5.6, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))
}
