package main

import (
	"fmt"
	"time"
)

func takesTime(name string, myChan chan<- string) {
	time.Sleep(time.Second * 2)
	fmt.Println("Hello : ", name)
	myChan <- "done"
}

func main() {
	start := time.Now()

	myChan := make(chan string, 6)
	// Channel can be buffered or unbuffered
	// Unbuffered channel has 0 capacity
	// myChan <- "Prayag" // deadlock since we have no receiver for unbuff

	go takesTime("Prayag", myChan)
	go takesTime("Bhatt", myChan)
	go takesTime("Prayag", myChan)
	go takesTime("Bhatt", myChan)
	go takesTime("Prayag", myChan)
	go takesTime("Bhatt", myChan)

	for len(myChan) < 2 {
	}

	fmt.Println(time.Since(start))
	close(myChan)
}
