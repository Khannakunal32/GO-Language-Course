package main

// run both files in this folder us "go run ."

import (
	"fmt"
	"strings"
)

func checkRemainingTickets (remainingTickets uint, orderedTicket uint) bool{

	if remainingTickets < orderedTicket {
		fmt.Printf("\nSorry, there are %v remaining tickets\n", remainingTickets)
		return false
	}
	return true
}

func printFirstNames(bookingsSlice []string) { // values: ["kunal khanna", "chhavi arora", "buddy khanna"] => ["kunal","chhavi","buddy"]

	var firstNameSlice [] string
	for _, nameElement := range bookingsSlice { // slice for loop give us index and an element to ignore the index or any variable we use _ so that it is ignored 
			
		var names = strings.Fields(nameElement); // "kunal khanna" => ["kunal","khanna"]
		firstNameSlice = append(firstNameSlice, names[0])
	}

	fmt.Printf("The first names of booking are %v\n", firstNameSlice)
}

func ValidateUserInput(firstName string, lastName string, email string) (bool, bool) {

	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")

	return isValidName, isValidEmail
} 