package main

import "fmt"

func main() {
	fmt.Println("Building StandAlones for Different Operating Systems !!")

	// in Terminal --> go env --> GOOS : "windows" or "linux" or "darwin" depends on type of OS
	// Using this we can build for Different OS regardless of where we are !!

	// GOOS="linux"  go build  --> to create StandAlone for Linux
	// GOOS="darwin" go build  --> to create StandAlone for MacOS

}
