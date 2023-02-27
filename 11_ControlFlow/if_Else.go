package main

import "fmt"

func main() {
	fmt.Println("If-Else Condition")

	var count int = 10

	if count2 := 11; count == 10 && count2 > 10 {
		fmt.Println("Equals")
	} else if count > 10 {
		fmt.Println("Count is greater than 10")
	} else {
		fmt.Println("Less than 10")
	}


}
