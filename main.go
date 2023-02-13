package main // packages to use

import (
	"fmt"
	"strings"
) // imported libraries
/// we use command like go mod init booking-app
// then to run we use go run filename

func main()  {
	var firstName string;
	var lastName string;
	var email string;
	var noOfTickets uint = 0;
	var remainingTickets uint= 50;
	var conferenceName string = "Go conference";
	var bookingsArray [50] string; // array of strings
	var bookingsSlice [] string; // slice of strings with not defined size
	
	
	// for{} runs loop infinite
	for checkRemainingTickets(remainingTickets, noOfTickets){
		fmt.Printf("\nPlease Enter your first name: ")
		fmt.Scan(&firstName);
	
		fmt.Printf("Please Enter your last name: ")
		fmt.Scan(&lastName);
		
		fmt.Printf("Please Enter your email: ")
		fmt.Scan(&email);
		
		fmt.Printf("Enter number of tickets to book: ")
		fmt.Scan(&noOfTickets);
	
		remainingTickets -= noOfTickets;
		bookingsArray[0] = firstName + " " + lastName;
		
		fmt.Printf("\nThe whole array: %v\n",bookingsArray)
		fmt.Printf("BookingsArray type: %T\n",bookingsArray)
		fmt.Printf("The length array: %v\n",len(bookingsArray))
		
		bookingsSlice = append(bookingsSlice, firstName+ " " + lastName) // to append to the last bookingSlice last
	
		fmt.Printf("\nThe whole Slice: %v\n",bookingsSlice)
		fmt.Printf("BookingsSlice type: %T\n",bookingsSlice)
		fmt.Printf("The length Slice: %v\n",len(bookingsSlice))
	
		fmt.Printf("\nThank you %v %v for booking %v tickets, \nYou will be notified on your email at %v\n", firstName, lastName, noOfTickets, email);
		fmt.Printf("%v remaing for %v\n", remainingTickets, conferenceName);
		
		//printout out the first names by creating a slice 

		
		
		// print first names
		printFirstNames(bookingsSlice);
	}
}

func checkRemainingTickets (remainingTickets uint, orderedTicket uint) bool{

	if remainingTickets < orderedTicket {
		fmt.Printf("\nSorry, there are %v remaining tickets\n", remainingTickets)
		return false
	}
	return true
}

func printFirstNames(bookingsSlice []string) {
	var firstNameSlice [] string
	for _, nameElement := range bookingsSlice { // slice for loop give us index and an element to ignore the index or any variable we use _ so that it is ignored 
			
		var names = strings.Fields(nameElement); // ["kunal khanna"] => [["kunal","khanna"]]
		firstNameSlice = append(firstNameSlice, names[0])
	}

	fmt.Printf("The first names of booking are %v\n", firstNameSlice)
}