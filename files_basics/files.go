package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error 

	newFile, err = os.Create("a.txt")

	if err != nil {
		log.Fatal(err)
	}
}