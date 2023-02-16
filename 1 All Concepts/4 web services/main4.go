package main

import (
	"errors"
	"fmt"
	"net/http"
)

// const portNumber = "127.0.0.1:8032"
const portNumber = ":8032"

//func for home route
func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome Home friends\n")
}

//func for home route
func aboutHandler(w http.ResponseWriter, r *http.Request){

	firstNumber := 100.0
	secondNumber := 100.0

	divident, _ := divisionOfTwoNumbers(firstNumber, secondNumber)

	fmt.Fprintf(w,"Divison of two numbers %v / %v: %v", firstNumber, secondNumber, divident)
}

// func to add sum of two number
func divisionOfTwoNumbers(a, b float64) (float64, error) {
	
	if b == 0 {
		err := errors.New("Cannot divide by zero")
	    return 0, err
	}

	return (a / b),nil
}


func main() {
	
	// creating route "/" and returning hello world 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Hello,world %v\n", r);
		n, err := fmt.Fprintf(w, "Hello,world Hi\n")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nvalue of n is %d\n", n)	
	})


	// creating route for home and passing home handler func
	http.HandleFunc("/home", homeHandler)
	
	// Same for about
   http.HandleFunc("/about", aboutHandler)

   // the below statement listen and serve for given port number but firewall pop up comes
	// http.ListenAndServe(":8032",nil)

	//this aviods any pop up from coming
	fmt.Printf("\n\nStarting server on port %s\n", portNumber)
	http.ListenAndServe(portNumber,nil)
}