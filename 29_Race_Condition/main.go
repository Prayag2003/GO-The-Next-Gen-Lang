package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions")

	waitGroup := &sync.WaitGroup{} // reference
	mut := &sync.Mutex{}           // reference

	var score = []int{0}

	waitGroup.Add(3)
	// IIFE
	go func(waitGroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine 1 is running")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		waitGroup.Done()
	}(waitGroup, mut)

	// waitGroup.Add(1)
	go func(waitGroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine 2 is running")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		waitGroup.Done()
	}(waitGroup, mut)

	// waitGroup.Add(1)
	go func(waitGroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine 3 is running")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		waitGroup.Done()
	}(waitGroup, mut)

	waitGroup.Wait()

	fmt.Println(score)
}
