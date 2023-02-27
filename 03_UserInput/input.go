package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome , Let's learn to take inputs ! \n"
	fmt.Println(welcome)

	// creating a reader using Bufio.Reader which reads Stdin from the OS Lib
	reader := bufio.NewReader(os.Stdin)

	// Since Go doesn't have any support for Try/Catch , we use Comma Ok / Error Ok Syntax
	// The Paradigm or language design of Golang is designed in such a way that it treats problems/errors in a true / false manner .

	// Comma - Ok    ||  Error - Ok
	// Try   - Catch ||  Try   - Catch

	fmt.Println("Enter Your Name : ")
	input, _ := reader.ReadString('\n') // reading will stop when reader finds \n
	fmt.Println("UserName is : ", input)

	fmt.Println("Enter Your Age : ")
	input2, _ := reader.ReadString('\n')
	fmt.Println("Age is : ", input2)

}
