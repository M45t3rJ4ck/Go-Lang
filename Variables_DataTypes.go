package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Variables
	fmt.Println("Variables in Go")
	var vName string = "Mickey"
	fmt.Println("Variable vName:", vName)
	// If use capital letter at the beginning of a variable name, it will be exported and can be used outside the package.
	var v1, v2 = 1.2, 3.4
	fmt.Println("Variable v1:", v1)
	fmt.Println("Variable v2:", v2)
	var v3 = "Hello"
	fmt.Println("Variable v3:", v3)
	v1 = 5.6
	fmt.Println("Variable v1:", v1)

	fmt.Println("")
	fmt.Println("Data Types in Go")
	//Data Types
	// int, float64, bool, string, runes
	fmt.Println("Data Type of 25:", reflect.TypeOf(25))
	fmt.Println("Data Type of 3.14:", reflect.TypeOf(3.14))
	fmt.Println("Data Type of true:", reflect.TypeOf(true))
	fmt.Println("Data Type of Mouse:", reflect.TypeOf("Mouse"))

	fmt.Println("")
	fmt.Println("Data Type Casting in Go")
	// If you want to use a variable without defining its type, you can use the short declaration operator :=.
	cV1 := 1.5
	cV2 := int(cV1)
	fmt.Println("Variable cV1:", cV1, "Data Type:", reflect.TypeOf(cV1))
	fmt.Println("Variable cV2:", cV2, "Data Type:", reflect.TypeOf(cV2))
	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)
	fmt.Println(cV4, err, "Data Type:", reflect.TypeOf(cV4))
	// strconv.Atoi converts a string to an int, returning an error if the conversion fails.
	if err == nil {
		fmt.Println("Variable cV3:", cV3, "Data Type:", reflect.TypeOf(cV3))
		fmt.Println("Variable cV4:", cV4, "Data Type:", reflect.TypeOf(cV4))
	} else {
		fmt.Println("Error converting cV3 to int:", err)
	}
	cV5 := 5000000
	cV6 := strconv.Itoa(cV5)
	// strconv.Itoa converts an int to a string.
	fmt.Println(cV6, err, "Data Type:", reflect.TypeOf(cV6))
	if err == nil {
		fmt.Println("Variable cV5:", cV5, "Data Type:", reflect.TypeOf(cV5))
		fmt.Println("Variable cV6:", cV6, "Data Type:", reflect.TypeOf(cV6))
	} else {
		fmt.Println("Error converting cV5 to string:", err)
	}
	cV7 := "3.14"
	// strconv.ParseFloat converts a string to a float64.
	// It returns the converted value and an error if the conversion fails.
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		fmt.Println("Variable cV7:", cV7, "Data Type:", reflect.TypeOf(cV7))
		fmt.Println("Variable cV8:", cV8, "Data Type:", reflect.TypeOf(cV8))
	} else {
		fmt.Println("Error converting cV7 to float64:", err)
	}
	// S Printf is used to format strings.
	cV9 := fmt.Sprintf("The value of cV8 is: %.2f", 3.14)
	fmt.Println("Variable cV9:", cV9, "Data Type:", reflect.TypeOf(cV9))
}
