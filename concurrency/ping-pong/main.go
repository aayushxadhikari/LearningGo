package main

import (
	"fmt"
	"sync"
)

func main() {
	const rounds = 5

	pingCh := make(chan bool)
	pongCh := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	// Ping Goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < rounds; i++ {
			<-pingCh
			fmt.Println("Ping")
			pongCh <- true
		}
	}()

	// Pong Goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < rounds; i++ {
			<-pongCh
			fmt.Println("Pong")
			if i < rounds-1 {
				pingCh <- true
			}
		}
	}()

	// Start the first Ping
	pingCh <- true

	wg.Wait()
	fmt.Println("Game over")
}
