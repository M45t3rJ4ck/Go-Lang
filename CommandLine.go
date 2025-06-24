package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	pl := fmt.Println

	pl(os.Args)

	args := os.Args[1:]

	iArgs := []int{}

	for _, i := range args {
		value, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		iArgs = append(iArgs, value)
	}

	max := 0

	for _, value := range iArgs {
		if value > max {
			max = value
		}
	}

	pl("The maximum value is:", max)
}
