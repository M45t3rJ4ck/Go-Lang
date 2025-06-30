package main

import (
	"fmt"
)

// var myApp map [keyType}valueType]
func main() {
	pf := fmt.Printf
	pl := fmt.Println

	var heroes map[string]string
	heroes = make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["Spiderman"] = "Peter Parker"

	villians := make(map[string]string)
	villians["Joker"] = "Unknown"
	villians["Lex Luthor"] = "Alexander Joseph Luthor"

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}

	pf("Batman is %v\n", heroes["Batman"])
	pf("Chip: " + superPets[3] + "\n")
	// Check if value at index
	_, ok := superPets[3]
	pl("Is there a 3rd pet:", ok)

	for key, value := range heroes {
		pf("%s is %s\n", key, value)
	}

	// Heroe goes insane
	delete(heroes, "Spiderman")
}
