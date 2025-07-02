package main

import (
	"fmt"
)

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	fmt.Println("A cat attacks it's prey.")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	fmt.Println("A cat says Hissssss.")
}

func (c Cat) HappySound() {
	fmt.Println("A cat says Purrrrrr.")
}

func main() {
	var kitty Animal

	kitty = Cat("Kitty")

	kitty.AngrySound()
	kitty.HappySound()

	// Type Converstion
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	fmt.Println("Cat's name:", kitty2.Name())
}
