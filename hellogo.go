package main

import (
	"bufio"
	"fmt"
	f "fmt"
	"log"
	"os"
)

var pl = f.Println

func main() {
	pl("Hello from Go")
	fmt.Print("What is your name?")
	// Buffered input
	reader := bufio.NewReader(os.Stdin)
	// Recieve text up to deffined point
	name, err := reader.ReadString('\n')
	// Catch error
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
		return
	}
}
