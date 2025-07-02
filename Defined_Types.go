package main

import (
	"fmt"
)

type Teaspoon float64
type Tablespoon float64
type Milliliters float64

func tspToMl(tsp Teaspoon) Milliliters {
	return Milliliters(tsp * 14.79)
}

func tbsToMl(tbs Tablespoon) Milliliters {
	return Milliliters(tbs * 4.92)
}

// Associated Methods
func (ml Milliliters) ToTsp() Teaspoon {
	return Teaspoon(ml / 14.79)
}

func (ml Milliliters) ToTsb() Tablespoon {
	return Tablespoon(ml / 4.92)
}

func main() {
	pf := fmt.Printf

	ml1 := Milliliters(6.9)

	pf("Teaspoon To Milliliters: %.2f ml\n", tspToMl(4.8))
	pf("Tablespoon To Milliliters: %.2f ml\n", tbsToMl(9.3))

	// Associated Methods
	pf("Milliliters To Teaspoon: %.2ftsp(s)\n", ml1.ToTsp())
	pf("Milliliters To Tablespoon: %.2ftbs(s)\n", ml1.ToTsb())
}
