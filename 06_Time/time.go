package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" Time Study Of Golang ")

	currTime := time.Now()
	// UnixNano returns the Number Of NanoSeconds elasped after 1970 ...
	currTimeInNano := time.Now().UnixNano()
	fmt.Println(currTime)
	fmt.Println(currTimeInNano)

	// using Standard Date for Formatting ( 01-02-2006 )*
	fmt.Println(currTime.Format("01-02-2006 Monday 15:04:05 "))

	// creating a DATE
	createdDate := time.Date(2003, time.March, 20, 12, 20, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))

}
