package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Hello,world %v\n", r);
		
		n, err := fmt.Fprintf(w, "Hello,world\n")
		
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nvalue of n is %d\n", n)

	})

	http.ListenAndServe(":8080",nil)
}