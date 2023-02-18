package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = "127.0.0.1:8080"
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	fmt.Println("Hello starting with the project........")
	log.Fatal(http.ListenAndServe(portNumber, nil))
}

