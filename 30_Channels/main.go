package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myChan := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// IMPORTANT: Receiver
	// Receiver  the values in the channel
	go func(ch <-chan int, wg *sync.WaitGroup) {

		// TODO: Check if the channel is open or not
		val, isChannelOpen := <-myChan

		fmt.Println(val)
		fmt.Println(isChannelOpen)

		// fmt.Println(<-myChan)
		// fmt.Println(<-myChan)
		wg.Done()
	}(myChan, wg)

	// Putting the values in the channel

	// IMPORTANT: Sender
	go func(ch chan<- int, wg *sync.WaitGroup) {

		myChan <- 0
		// myChan <- 6
		close(myChan)

		wg.Done()
	}(myChan, wg)

	wg.Wait()

	// fmt.Println(<-myChan)
	// myChan <- 5
}
