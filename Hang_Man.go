package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var hmArr = [7]string{
	"  +---+\n" +
		"      |\n" +
		"      |\n" +
		"      |\n" +
		"=========\n",
	"  +---+\n" +
		"  0   |\n" +
		"      |\n" +
		"      |\n" +
		"=========\n",
	"  +---+\n" +
		"  0   |\n" +
		"  |   |\n" +
		"      |\n" +
		"=========\n",
	"  +---+\n" +
		"  0   |\n" +
		" /|   |\n" +
		"      |\n" +
		"=========\n",
	"  +---+\n" +
		"  0   |\n" +
		" /|\\  |\n" +
		"      |\n" +
		"=========\n",
	"  +---+\n" +
		"  0   |\n" +
		" /|\\  |\n" +
		" /    |\n" +
		"=========\n",
	"  +---+\n" +
		"  0   |\n" +
		" /|\\  |\n" +
		" / \\  |\n" +
		"=========\n",
}

var wordArr = [7]string{
	"zodiac",
	"zipzag",
	"zipper",
	"orange",
	"zombie",
	"kiwi",
	"jazz",
}

var randWord string
var userGuess string
var guessedLetters string
var wrongGuesses []string
var correctLetters []string

func getRandWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randWord := wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))

	return randWord
}

func showBoard() {
	fmt.Println(hmArr[len(wrongGuesses)])
	fmt.Print("Secret word: ")

	for _, value := range correctLetters {
		if value == "" {
			fmt.Print("_")
		} else {
			fmt.Print(value)
		}
	}

	fmt.Print("\nIncorrect guesses: ")

	if len(wrongGuesses) > 0 {
		for _, value := range wrongGuesses {
			fmt.Print(value + " ")
		}
	}

	fmt.Println()
}

func getUserLetter() string {
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("Please enter a letter: ")

		userGuess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		userGuess = strings.ToUpper(userGuess)
		userGuess = strings.TrimSpace(userGuess)

		isLetter := regexp.MustCompile(`^[a-zA-Z]$`).MatchString

		if len(userGuess) != 1 {
			fmt.Println("Please enter only one letter.")
		} else if !isLetter(userGuess) {
			fmt.Println("Please enter a letter: ")
		} else if strings.Contains(guessedLetters, userGuess) {
			fmt.Println("Please enter a letter you haven't guessed yet.")
		} else {
			return userGuess
		}
	}

	return userGuess
}

func getAllIndexes(theStr, subStr string) (indices []int) {
	if (len(subStr) == 0) || (len(theStr) == 0) {
		return indices
	}

	offset := 0

	for {
		i := strings.Index(theStr[offset:], subStr)

		if i == -1 {
			return indices
		}

		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func updateCorrectLetters(letter string) {
	indexMatches := getAllIndexes(randWord, letter)

	for _, value := range indexMatches {
		correctLetters[value] = letter
	}
}

func sliceHasEmptys(theSlice []string) bool {
	for _, value := range theSlice {
		if len(value) == 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(getRandWord())
	fmt.Println("Welcome to Hangman!")
	fmt.Println("You have 6 incorrect guesses before you lose.")

	for true {
		showBoard()

		userGuess := getUserLetter()

		if strings.Contains(randWord, userGuess) {
			updateCorrectLetters(userGuess)

			if sliceHasEmptys(correctLetters) {
				fmt.Println("You guessed a letter that is in the word!")
			} else {
				fmt.Println("Congratulations! You guessed the word: ", randWord)
				break
			}
		} else {
			guessedLetters += userGuess
			wrongGuesses = append(wrongGuesses, userGuess)

			if len(wrongGuesses) >= 6 {
				fmt.Println("You have run out of guesses! The word was: ", randWord)
				break
			}
		}
	}
}
