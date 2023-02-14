package main

import (
	"log"
	"overview-app/helper"
)

const numPool = 1000

func CalculateValue(intchan chan int){
	randomNumber := helper.RandomNumber(numPool)

	intchan <- randomNumber // this says push the value of random num into intChan channel
}

func main() {
	intChan := make(chan int)
	defer close(intChan)  // defer says that whatever comes after this keyword execute it as soon as this current function overs

	go CalculateValue(intChan) // to use channel we make this a routine of itself using go keyword without keyword this will not work
	
	num := <-intChan
	log.Println(num)
}