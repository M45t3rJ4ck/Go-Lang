package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Basic Math:")
	fmt.Println("5 + 4 = ", 5+4)
	fmt.Println("5 - 4 = ", 5-4)
	fmt.Println("5 * 4 = ", 5*4)
	fmt.Println("5 / 4 = ", 5/4)
	fmt.Println("5 % 4 = ", 5%4)

	fmt.Println("")
	fmt.Println("Increment and Decrement:")
	mInt := 5
	fmt.Println("mInt = ", mInt)
	mInt++ // Increment
	fmt.Println("mInt++ = ", mInt)
	mInt-- // Decrement
	fmt.Println("mInt-- = ", mInt)

	fmt.Println("")
	fmt.Println("Float precision:")
	fmt.Println("Float Precision of 3.14159265358979323846 + 3.14159265358979323846: ", 3.14159265358979323846+3.14159265358979323846)

	fmt.Println("")
	fmt.Println("Generating Random Numbers:")
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(100) + 1 // Random number between 0 and 100
	fmt.Println("Random number between 1 and 100: ", randNum)

	fmt.Println("")
	fmt.Println("Common Math Functions:")
	fmt.Println("Absolute value of -42: ", math.Abs(-42))
	fmt.Println("Power of 2^3: ", math.Pow(2, 3))
	fmt.Println("Square root of 16: ", math.Sqrt(16))
	fmt.Println("Cube root of 27: ", math.Cbrt(27))
	fmt.Println("Ceiling of 4.2: ", math.Ceil(4.2))
	fmt.Println("Floor of 4.8: ", math.Floor(4.8))
	fmt.Println("Rounding 4.5: ", math.Round(4.5))
	fmt.Println("Logarithm base 2 of 100: ", math.Log2(100))
	fmt.Println("Logarithm base 10 of 100: ", math.Log10(100))
	fmt.Println("Maximum of 5 and 10: ", math.Max(5, 10))
	fmt.Println("Minimum of 5 and 10: ", math.Min(5, 10))
	fmt.Println("Trigonometric sine of 90 degrees: ", math.Sin(math.Pi/2))
	fmt.Println("Trigonometric cosine of 0 degrees: ", math.Cos(0))
	fmt.Println("Trigonometric tangent of 45 degrees: ", math.Tan(math.Pi/4))

	fmt.Println("")
	fmt.Println("Converting Degrees to Radians and Vice Versa:")
	r90 := 90 * math.Pi / 180 // Convert degrees to radians
	fmt.Println("90 degrees in radians: ", r90)
	d90 := r90 * (180 / math.Pi)
	fmt.Println("90 radians in degrees: ", d90)
}
