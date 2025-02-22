// Buffered Channels
package main

import(
	"fmt"
)

func sendData2(ch chan int){
	for i:= 0; i<3; i++{
		ch <- i
		fmt.Println("Sent:",i)
	}
} 

func main2(){
	ch:= make(chan int, 3)

	go sendData2(ch)

	for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Println("Received:", value)
	}
}