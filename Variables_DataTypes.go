package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	// Variables
	pl("Variables in Go")
	var vName string = "Mickey"
	pl("Variable vName:", vName)
	// If use capital letter at the beginning of a variable name, it will be exported and can be used outside the package.
	var v1, v2 = 1.2, 3.4
	pl("Variable v1:", v1)
	pl("Variable v2:", v2)
	var v3 = "Hello"
	pl("Variable v3:", v3)
	v1 = 5.6
	pl("Variable v1:", v1)

	pl("")
	pl("Data Types in Go")
	//Data Types
	// int, float64, bool, string, runes
	pl("Data Type of 25:", reflect.TypeOf(25))
	pl("Data Type of 3.14:", reflect.TypeOf(3.14))
	pl("Data Type of true:", reflect.TypeOf(true))
	pl("Data Type of Mouse:", reflect.TypeOf("Mouse"))

	pl("")
	pl("Data Type Casting in Go")
	// If you want to use a variable without defining its type, you can use the short declaration operator :=.
	cV1 := 1.5
	cV2 := int(cV1)
	pl("Variable cV1:", cV1, "Data Type:", reflect.TypeOf(cV1))
	pl("Variable cV2:", cV2, "Data Type:", reflect.TypeOf(cV2))
	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, "Data Type:", reflect.TypeOf(cV4))
	// strconv.Atoi converts a string to an int, returning an error if the conversion fails.
	if err == nil {
		pl("Variable cV3:", cV3, "Data Type:", reflect.TypeOf(cV3))
		pl("Variable cV4:", cV4, "Data Type:", reflect.TypeOf(cV4))
	} else {
		pl("Error converting cV3 to int:", err)
	}
	cV5 := 5000000
	cV6 := strconv.Itoa(cV5)
	// strconv.Itoa converts an int to a string.
	pl(cV6, err, "Data Type:", reflect.TypeOf(cV6))
	if err == nil {
		pl("Variable cV5:", cV5, "Data Type:", reflect.TypeOf(cV5))
		pl("Variable cV6:", cV6, "Data Type:", reflect.TypeOf(cV6))
	} else {
		pl("Error converting cV5 to string:", err)
	}
	cV7 := "3.14"
	// strconv.ParseFloat converts a string to a float64.
	// It returns the converted value and an error if the conversion fails.
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl("Variable cV7:", cV7, "Data Type:", reflect.TypeOf(cV7))
		pl("Variable cV8:", cV8, "Data Type:", reflect.TypeOf(cV8))
	} else {
		pl("Error converting cV7 to float64:", err)
	}
	// S Printf is used to format strings.
	cV9 := fmt.Sprintf("The value of cV8 is: %.2f", 3.14)
	pl("Variable cV9:", cV9, "Data Type:", reflect.TypeOf(cV9))
}
