package main

import "log"

func main1(){
	var myString string 
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("myString is change to", myString)
}

func changeUsingPointer(s *string){
	newValue := "Red"
	*s = newValue
}