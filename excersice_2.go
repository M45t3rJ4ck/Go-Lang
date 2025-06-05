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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your first number: ")
	firstNum, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	firstNum = strings.TrimSpace(firstNum)
	iFirstNum, err := strconv.Atoi(firstNum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter your second number: ")
	secondNum, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	secondNum = strings.TrimSpace(secondNum)
	iSecondNum, err := strconv.Atoi(secondNum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter your logic operator: ")
	logicOperator, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	logicOperator = strings.TrimSpace(logicOperator)

	if logicOperator == "+" {
		fmt.Printf("%s + %s = %d\n", firstNum, secondNum, (iFirstNum + iSecondNum))
	} else if logicOperator == "-" {
		fmt.Printf("%s - %s = %d\n", firstNum, secondNum, (iFirstNum - iSecondNum))
	} else if logicOperator == "*" {
		fmt.Printf("%s * %s = %d\n", firstNum, secondNum, (iFirstNum * iSecondNum))
	} else if logicOperator == "/" {
		fmt.Printf("%s / %s = %d\n", firstNum, secondNum, (iFirstNum / iSecondNum))
	} else if logicOperator == "%" {
		fmt.Printf("%s %% %s = %d\n", firstNum, secondNum, (iFirstNum % iSecondNum))
	} else {
		log.Fatal(err)
	}
}
