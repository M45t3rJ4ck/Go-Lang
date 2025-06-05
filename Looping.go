package main

import (
	"fmt"
)

func main() {
	fmt.Println("Looping in Go")
	// Standard for loop
	// For initialization; condition: postStatement {BODY}
	for x := 1; x <= 5; x++ {
		fmt.Println("x added is: ", x)
	}
	for x := 5; x >= 1; x-- {
		fmt.Println("x minused is:", x)
	}
	//For mimic while loop
	fX := 1
	for fX <= 5 {
		fmt.Println("fX added is: ", fX)
		fX++
	}

	//Array with numbers
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		fmt.Println("Number in array is: ", num)
	}
	//For loop mimic do while loop
	xV1 := 1
	for true {
		if xV1 == 5 {
			fmt.Println("Break")
			break
		}
		fmt.Println("xV1 is: ", xV1)
		xV1++
	}
}
