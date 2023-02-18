package main

import (
	"fmt"
	"log"
	"net/http"
	
)


func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	fmt.Println("Hello starting with the project........")
	log.Fatal(http.ListenAndServe(portNumber, nil))
}

