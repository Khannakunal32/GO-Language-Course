package main

import (
	"fmt"
)

// struct with functions
type person struct {
	firstName string
}

func (pObject *person) printFirstName() string {
	return pObject.firstName
}

func main() {
	var structVar1 person
	structVar1.firstName = "chhavi"

	var structVar2 = person{
		firstName: "Kunal",
	}

	fmt.Printf("\nPerson first name is: %v\n",structVar2.printFirstName())

}