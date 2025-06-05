package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Welcome, What is your age? ")
	reader := bufio.NewReader(os.Stdin)
	// Recieve text up to deffined point
	age, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	age = strings.TrimSpace(age)
	userAge, err := strconv.Atoi(age) // Convert string to int, removing the newline character

	if err != nil {
		log.Fatal(err)
	}

	if userAge < 5 {
		fmt.Print("Your to young, stay at home.")
	} else if userAge == 6 {
		fmt.Print("Your old enough for kindergarden.")
	} else if (userAge >= 6) || (userAge <= 18) {
		grade := userAge - 6
		fmt.Print("Your old enough to go to grade ", grade, ".")
	} else {
		fmt.Print("Your old enough for college.")
	}
}
