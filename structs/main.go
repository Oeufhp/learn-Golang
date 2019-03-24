package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Inwza",
		contact: contactInfo{
			email:   "jim@eiei.com",
			zipCode: 10510,
		},
	}
	fmt.Printf("%+v", jim)
}
