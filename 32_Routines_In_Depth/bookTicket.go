package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ticketsAvailable = 5
	mu               sync.Mutex
)

func bookTicket() {
	mu.Lock()
	defer mu.Unlock()

	if ticketsAvailable > 0 {
		time.Sleep(100 * time.Millisecond)

		// Book a ticket
		ticketsAvailable--
		fmt.Printf("Ticket booked. %d ticket(s) remaining.\n", ticketsAvailable)
	} else {
		fmt.Println("Sorry, no more tickets available.")
	}
}

func main() {
	var wg sync.WaitGroup
	numThreads := 5

	for i := 1; i <= numThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Thread %d trying to book a ticket...\n", id)
			bookTicket()
		}(i)
	}

	wg.Wait()
}
