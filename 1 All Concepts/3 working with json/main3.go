// working with JSON
package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"` // tells go that we are creating json what it should look for
	LastName string `json:"last_name"`
	Age int `json:"age"`
	HasDog bool `json:"has_dog"`
}


func main() {
	myJson := `
	[
		{
			"first_name": "kunal",
			"last_name": "Khanna",
			"age": 18,
			"has_dog": true
		},
		{
			"first_name": "Chhavi",
			"last_name": "Arora",
			"age": 20,
			"has_dog": false
		}
	]`

	// unmarshalling ie extracting data from json

	var unmarshalled []Person //slice of struct
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("error while unmarshalling",err)
	}
	
	log.Printf("\nunmarshalled data: %v\n",unmarshalled)
	
	
	// marshalling
	person1 := Person{
		FirstName: "Buddy",
		LastName: "Khanna",
		Age: 5,
		HasDog: true,
	}
	
	person2 := Person{
		FirstName: "Scotty",
		LastName: "Khanna",
		Age: 10,
		HasDog: true,
	}

	var myStructSlice []Person
	myStructSlice = append(myStructSlice, person1)
	myStructSlice = append(myStructSlice, person2)

	marshalled , err := json.MarshalIndent(myStructSlice,"","     ")
	if err!=nil {
		log.Println("error while marshalling",err)
	}

	// the marshalled format must be stringified
	log.Printf("\nMarshalled data from struct is: %v\n", string(marshalled))
	
}