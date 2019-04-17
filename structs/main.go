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

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Inwza",
		contactInfo: contactInfo{
			email:   "jim@eiei.com",
			zipCode: 10510,
		},
	}

	// jimPointer := &jim
	jim.updateName("eiei")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerPerson *person) updateName(newFirstname string) {
	(*pointerPerson).firstName = newFirstname
}
