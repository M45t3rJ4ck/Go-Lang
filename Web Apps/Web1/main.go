package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nb, err := fmt.Fprintf(w, "Hello Browser")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes Written: %d\n", nb)
	})
	err := http.ListenAndServe("localhost: 5000", nil)
	log.Fatal(err)
}
