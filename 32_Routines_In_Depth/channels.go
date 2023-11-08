package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Create a buffered channel with a capacity of 2
	messageChannel := make(chan string, 2)
	wg.Add(2)

	// First goroutine sends a message
	go func() {
		fmt.Println("")
		defer wg.Done()
		messageChannel <- "Hello from Goroutine 1!"
		fmt.Println("Goroutine 1 sent a message.")
	}()

	// Second goroutine receives and prints the message
	go func() {
		defer wg.Done()
		message := <-messageChannel
		fmt.Println("Goroutine 2 received:", message)
	}()

	wg.Wait()
}
