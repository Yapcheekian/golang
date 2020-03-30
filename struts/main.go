package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastname  string
	contactInfo
}

func main() {
	person := person{
		firstName: "Alex",
		lastname:  "Anderson",
		contactInfo: contactInfo{
			email:   "gmail.com",
			zipCode: 12345,
		},
	}
	personPointer := &person
	personPointer.updateName("yap")
	person.print()
}

func (pointerToPerson *person) updateName(newFirstname string) {
	(*pointerToPerson).firstName = newFirstname
}

func (p person) print() {
	fmt.Println(p)
}
