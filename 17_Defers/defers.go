package main

import "fmt"

func main() {

	fmt.Println("Defers in Golang ... ")

	// Defer sends the Function to the last line above curly braces and if there are multiple defers then they will be send in FIFO Order

	defer fmt.Println("23") // deferred to the end
	defer fmt.Println("24") // deferred to the end
	defer fmt.Println("25") // deferred to the end

	fmt.Println("Hello")
	myDefer()

}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
