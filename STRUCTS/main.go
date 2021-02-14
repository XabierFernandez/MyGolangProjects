package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	xabier := person{
		firstname: "Xabier",
		lastname:  "Fernandez",
		contactInfo: contactInfo{
			email:   " xabfer@info.com",
			zipCode: 99509,
		},
	}
	xabier.updateName("Xabi")
	xabier.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstname = newFirstName
}

func (p person) updateLastName(newLastName string) {
	p.lastname = newLastName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
