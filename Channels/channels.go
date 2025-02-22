// Basic example of channel
package main

import (
	"fmt"
)

func sendData(ch chan int) {
	ch <- 42 // Send the value 42 into the channel
}

func main3() {
	ch := make(chan int) // Create a new channel of type int

	go sendData(ch) // Run sendData in a goroutine

	value := <-ch // Receive the value from the channel
	fmt.Println("Received:", value) // Output the received value
}
