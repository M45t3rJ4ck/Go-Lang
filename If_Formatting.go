package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional Operators in Go")
	// If Conditionals
	// Conditional Operators: ==, !=, <, <=, >, >=
	// Logical Operators: &&, ||, !
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		fmt.Println("This is an Important Birthday!")
	} else if (iAge == 21) || (iAge == 50) {
		fmt.Println("This is an Important Birthday Too!")
	} else if iAge >= 65 {
		fmt.Println("This is an Important Birthday Too!")
	} else {
		fmt.Println("This is not an Important Birthday")
	}
	fmt.Println("!true = ", !true)
	fmt.Println("!false = ", !false)

	fmt.Println("")
	fmt.Println("Formatted Output in Go")
	// Formatted Output
	/*
		%d - integer
		%c - character
		%f - float
		%t - boolean
		%s - string
		%o - Base 8
		%x - Base 16
		%v - default format
		%T - type of value
		%q - quoted string
		%p - pointer
	*/
	fmt.Printf("%s %d %c %f %t %o %x\n", "Hello", 42, 'A', 3.14, true, 42, 42)
	fmt.Printf("%9f\n", 3.14159265358979323846)        // 9 characters wide, right-aligned
	fmt.Printf("%.2f\n", 3.14159265358979323846)       // 2 decimal places
	fmt.Printf("%9.f\n", 3.14159265358979323846)       // 9 characters wide, no decimal places
	sp1 := fmt.Sprintf("%9.f", 3.14159265358979323846) // formatted string
	fmt.Println("Formatted String: ", sp1)
}
