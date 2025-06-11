package main

import (
	"fmt"
)

// func functionName(parameters) returnType {BODY}

func sayHello() {
	fmt.Println("Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("Can not divide by zero")
	} else {
		return x / y, nil
	}
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func changeValue(f3 int) int {
	f3 += 1
	return f3
}

func main() {
	sayHello()

	fmt.Println(getSum(5, 6))

	ans, err := getQuotient(5, 0)
	if err == nil {
		fmt.Printf("%f / %f = %f", 5.0, 0.0, ans)
	} else {
		fmt.Println(err)
	}

	f1, f2 := getTwo(5)
	fmt.Printf("%d %d\n", f1, f2)

	fmt.Println("Unknown Sum: ", getSum2(1, 2, 3, 4, 5))

	vArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array Sum: ", getArraySum(vArr))

	f3 := 5
	fmt.Println("F3 before funnction: ", f3)
	changeValue(f3)
	fmt.Println("F3 after function: ", f3)
}
