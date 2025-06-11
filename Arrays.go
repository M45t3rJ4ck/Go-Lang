package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	fmt.Println("Third index in arr1:", arr1[3])
	fmt.Println("Array length of arr1:", len(arr1))
	for i := 0; i < len(arr1); i++ {
		fmt.Println("Index", i, "of arr1:", arr1[i])
	}

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Second index in arr2:", arr2[2])
	fmt.Println("Array length of arr2:", len(arr2))
	for index, value := range arr2 {
		fmt.Printf("Index number %d in arr2: %d\n", index, value)
	}

	// Multi-dimensional array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr3[i][j])
		}
	}

	// Slicing runes
	aStr1 := "abcdefg"
	rArr1 := []rune(aStr1)
	for _, value := range rArr1 {
		fmt.Printf("Rune Array: %d\n", value)
	}
	byteArr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	bStr := string(byteArr[:])
	fmt.Println("I'm a string: ", bStr)
}
