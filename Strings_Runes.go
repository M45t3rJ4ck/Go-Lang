package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Example of using strings in Go:
	sV1 := "A Word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)

	fmt.Println(sV1)
	fmt.Println(sV2)
	fmt.Println("Length of sV1: ", len(sV1))
	fmt.Println("Length of sV2: ", len(sV2))
	fmt.Println("Contains the word 'another': ", strings.Contains(sV2, "Another"))
	fmt.Println("Index of letter 'o' in sV1: ", strings.Index(sV1, "o"))
	fmt.Println("Index of letter 'o' in sV2: ", strings.Index(sV2, "o"))
	fmt.Println("Replace all instances of 'o': ", strings.Replace(sV1, "o", "0", -1))
	fmt.Println("Replace all instances of 'o': ", strings.Replace(sV2, "o", "0", -1))

	sV3 := "Here are\n some words\n to use"
	sV3 = strings.TrimSpace(sV3)
	fmt.Println("Original sV3: ", sV3)
	fmt.Println("Trimmed sV3: ", sV3)
	fmt.Println("Split sV3 by spaces: ", strings.Split(sV3, " "))
	fmt.Println("Lowercase of sV3: ", strings.ToLower(sV3))
	fmt.Println("Uppercase of sV3: ", strings.ToUpper(sV3))
	fmt.Println("sV3 has prefix 'Here': ", strings.HasPrefix(sV3, "Here"))
	fmt.Println("sV3 has suffix 'use': ", strings.HasSuffix(sV3, "use"))

	// Example of using runes in Go:
	rStr := "abcdefg"
	fmt.Println("Original string: ", rStr)
	fmt.Println("Rune Count: ", utf8.RuneCountInString(rStr))

	for i, runeValue := range rStr {
		fmt.Printf("Index: %d, Unicode: %#U, Rune: %c\n", i, runeValue, runeValue)
	}
}
