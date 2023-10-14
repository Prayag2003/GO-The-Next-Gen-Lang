package main

import "fmt"

func main() {

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("")
	// i returns the index not the value
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("")

	for index, day := range days {
		fmt.Printf("Day %v is %v \n", index+1, day)
	}

	index := 0
	for index < 10 {

		if index == 7 {
			goto prayag
		}
		if index == 8 {
			break
		}

		if index == 5 {
			index++
			continue
		}
		fmt.Println(index)
		index++
	}

prayag:
	fmt.Printf("You are exited ... ")
}
