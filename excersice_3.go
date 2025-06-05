package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	randNum := rand.Intn(50) + 1

	for true {
		fmt.Print("Guess a number between 1 and 50: ")
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)

		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randNum {
			fmt.Println("Your number was too high, try again: ")
		} else if iGuess < randNum {
			fmt.Println("Your number was too low, try again: ")
		} else {
			fmt.Println("You guessed the number!")
			break
		}
	}
}
