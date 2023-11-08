package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go runRoutine()
	wg.Add(1)
	fmt.Println("Hello Main")
	// time.Sleep(time.Second)
	wg.Wait()
}
func runRoutine() {
	fmt.Println("Hello Routine")
	wg.Done()
}
