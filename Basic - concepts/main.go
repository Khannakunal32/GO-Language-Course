package main // packages to use

import (
	kali "booking-app/greetings" // we need to import our made package into the by moduleName/functionName
	"fmt"
)

// imported libraries
/// we use command like go mod init booking-app
// then to run we use go run filename

// global variables can be defined here
// not importing fuctnios from same level is direct and does not require any importing or exporting but for another level ie package formed inside another folder require it

func main()  {
	var firstName string;
	var lastName string;
	var email string;
	var noOfTickets uint = 0;
	var remainingTickets uint= 50;
	var conferenceName string = "Go conference";
	var bookingsArray [50]string; // array of strings
	var bookingsSlice []string; // slice of strings with not defined size
	var bookingsMap []map[string]string // slice of bookingsMap
	var newInformationString = "true"
	
	kali.Greetings() // kali function is imported from the package named greetings 

	// for{} runs loop infinite
	for newInformationString == "true" && checkRemainingTickets(remainingTickets, noOfTickets) {
		fmt.Printf("\nPlease Enter your first name: ")
		fmt.Scan(&firstName);
	
		fmt.Printf("Please Enter your last name: ")
		fmt.Scan(&lastName);
		
		fmt.Printf("Please Enter your email: ")
		fmt.Scan(&email);
		
		fmt.Printf("Enter number of tickets to book: ")
		fmt.Scan(&noOfTickets);
		
		isValidName, isValidEmail := ValidateUserInput(firstName, lastName, email)
		
		if isValidEmail && isValidName {
			
			
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
			fmt.Printf("%v remaing for %v\n", remainingTickets, conferenceName)
						
			// print first names
			printFirstNames(bookingsSlice)

			// storing the map 
			var mapElement map[string]string = workingWithMap(firstName, lastName, email, noOfTickets);
			
			// adding the map element created in slice of map
			bookingsMap = append(bookingsMap, mapElement)

			// storing infromation in struct
			var userDataStruct = pushDataIntoStruct(firstName, lastName, email, noOfTickets)

			fmt.Printf("\nShowing data in struct: %v, like first name is %v\n", userDataStruct, userDataStruct.firstName)

			} else {

				fmt.Printf("\nNot valid inputs\n")
			}
			
			fmt.Println("You want to continue adding ?")
			fmt.Scan(&newInformationString);
			
		}
		//showing the information store in map
		printMapInformation(bookingsMap)

}
