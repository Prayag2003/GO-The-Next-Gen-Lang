// https://pkg.go.dev/runtime
package main

import "fmt"

func main() {
	fmt.Println("Managing Memory in GO !!! ")
	fmt.Println("Memory Allocation and Dellocation happens automatically in Golang !! ")

	//        N E W                  |              M A K E
	// ______________________________|__________________________________
	// 1) memory allocated but       |   1) memory allocated and
	//  zero-initialized              |		initialized based on the type
	//                               |
	// 2) we will get a memory       |   2) we will get a memory
	//    address                    |      address
	//                               |
	// 3) zero storage               |   3) Non-Zero Storage

	// Anything which is Out Of Scope or nil is eligible for Garbage Collection ..... After a certain ThresHold is reached , it starts collecting Garbage
}
