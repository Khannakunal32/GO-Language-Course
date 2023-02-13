package main // packages to use

import "fmt" // imported libraries

/// we use command like go mod init booking-app
// then to run we use go run filename

func main()  {
	var firstName string;
	var lastName string;
	var email string;
	var noOfTickets uint;
	var remainingTickets uint= 50;
	var conferenceName string = "Go conference";
	var bookingsArray [50] string; // array of strings
	var bookingsSlice [] string; // slice of strings with not defined size
	
	fmt.Printf("Please Enter your first name: ")
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

	
	// for loop infinite times


}