package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	pl := fmt.Println

	file, err := os.Create("fileio_test.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	iPrimeArr := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	var sPrimeArr []string

	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	for _, num := range sPrimeArr {
		_, err := file.WriteString(num + "\n")

		if err != nil {
			log.Fatal("Error writing to file:", err)
		}
	}

	file, err = os.Open("fileio_test.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scan1 := bufio.NewScanner(file)
	for scan1.Scan() {
		pl("Prime: ", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	_, err = os.Stat("fileio_test.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File does not exist.")
	} else {
		file, err = os.OpenFile("fileio_test.txt", os.O_APPEND|os.O_WRONLY|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Error opening file:", err)
		}
		defer file.Close()
		if _, err := file.WriteString("56\n"); err != nil {
			log.Fatal("Error writing to file:", err)
		}
	}
}
