package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	//someFunc("1") // this function is being run synchronously
	go someFunc("1") // this is known as the Go Routines this is asynchronously being called
	go someFunc("2")
	go someFunc("3")

	// without this line the main function exists immediately before receiving the data from the functions
	time.Sleep(time.Second * 2)


	fmt.Println(("hi"))
}