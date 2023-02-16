package main

import (
	"fmt"
	"net/http"
)

const portNumber = "127.0.0.1:8032"

//func for home route
func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome Home\n")
}

//func for home route
func aboutHandler(w http.ResponseWriter, r *http.Request){

	firstNumber := 10
	secondNumber := 20
	sum, _ := sumOfTwoNumbers(firstNumber, secondNumber)

	fmt.Fprintf(w,"\nThis is welcome page: \nSum of two numbers %v + %v: %v", firstNumber, secondNumber, sum)
}

// func to add sum of two number
func sumOfTwoNumbers(a, b int) (int, error) {
	
	return (a + b),nil
}


func main() {
	
	// creating route "/" and returning hello world 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Hello,world %v\n", r);
		n, err := fmt.Fprintf(w, "Hello,world\n")
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
	http.ListenAndServe(portNumber,nil)
}