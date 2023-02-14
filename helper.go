package main

// run both files in this folder us "go run ."

import (
	"fmt"
	"strconv"
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

// to store data in map we can use like key pair
// name: "Kunal Khanna" 

func workingWithMap(firstName string, lastName string, email string, userTickets uint) map[string]string{

	mapDataArray := make(map[string]string); // format is map[keyType]string, not the value of value is string type we cannot do anything to change it
	
	mapDataArray["firstName"] = firstName;
	mapDataArray["LastName"] = firstName;
	mapDataArray["email"] = firstName;
	mapDataArray["userTickets"] = strconv.FormatUint(uint64(userTickets), 10); // method to store intiger in a string format

	return mapDataArray
}

func printMapInformation(bookingsMap []map[string]string){
	
	for _, bookings := range bookingsMap {

		fmt.Printf("\nList of books from bookingsMap: %v\n", bookings)
	}
}