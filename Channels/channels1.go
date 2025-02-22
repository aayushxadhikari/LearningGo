// Channel Synchronize goroutines 
package main

import(
	"fmt"
	"time"
)

func sendData1(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "Hello from Goroutine"
}

func main1(){               // Remove the 1 before running the program
	ch := make(chan string)

	go sendData1(ch) // Send data in the goroutine

	fmt.Println("Waiting for data.....")
	msg := <- ch // Blocked until data is received
	fmt.Println("Received:", msg)


}