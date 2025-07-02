package main

import (
	"fmt"
)

type customer struct {
	name    string
	address string
	balance float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us $%.2f.\n", c.name, c.balance)
}

func newCustAdd(c *customer, address string) {
	c.address = address
	fmt.Printf("%s address changed to %s.\n", c.name, address)
}

// Area calculation
type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

// Compatition
type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact %s %s at %s on %s for more information.", b.contact.fName, b.contact.lName, b.name, b.contact.phone)
}

func main() {
	var tS customer

	tS.name = "Tom Smith"
	tS.address = "5 Main Street"
	tS.balance = 234.56

	getCustInfo(tS)
	newCustAdd(&tS, "123 Side Road")

	// Struct literral
	sS := customer{"Sally Smith", "123 Main Street", 0.00}
	getCustInfo(sS)

	// Area calculation
	rect1 := rectangle{10.5, 11.6}
	fmt.Println("Reactangle Area:", rect1.Area())

	// Compatition
	con1 := contact{
		"James", "Wang", "555-1212",
	}
	bus1 := business{
		"ABC Plumbing",
		"234 Side Road",
		con1,
	}
	bus1.info()
}
