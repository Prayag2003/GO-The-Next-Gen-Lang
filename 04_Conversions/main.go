package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Enter user input: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	// Converting to integer / float
	// StringtoInt, err := strconv.ParseFloat(input, 64)
	StringtoInt, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error")
		panic(err)
	} else {
		fmt.Println(StringtoInt + 1)
	}

}
