package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./Templates/" + tmpl)
	errorCheck(err)
	err = parsedTemplate.Execute(w, nil)
	errorCheck(err)
}

func homeHandler(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func aboutHandler(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func addHandler(w http.ResponseWriter, request *http.Request) {
	write(w, "The answer to the sum is:\n")
	sum := getSum(5, 10)
	output := fmt.Sprintf("\t5 + 10 = %d\n", sum)
	write(w, output)
}

func getSum(x, y int) int {
	return x + y
}

func divideHandler(w http.ResponseWriter, request *http.Request) {
	write(w, "The answer to the division is:\n")
	divide, err := getQuotient(5, 0)
	if err != nil {
		write(w, "\tCan not divide by zero.\n")
	} else {
		output := fmt.Sprintf("\t5 / 0 = %.2f\n", divide)
		write(w, output)
	}
}

func getQuotient(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Can not divide by zero.")
		return 0, err
	} else {
		return (x / y), nil
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/getSum", addHandler)
	http.HandleFunc("/getQuotient", divideHandler)

	err := http.ListenAndServe(":5000", nil)
	errorCheck(err)
}
