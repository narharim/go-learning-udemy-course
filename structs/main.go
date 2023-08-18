package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	bob := person{
		firstName: "bob",
		lastName:  "jelly",
		contactInfo: contactInfo{
			email:   "bob@gmail.com",
			zipCode: 749095,
		},
	}

	bob.updateName("Alice")
	bob.print()
}
