package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with Files in GO !!! ")

	// adding content to a file
	content := "This is my Content and let's add to the File !!! \n"

	// creating a File by using OS Library in the Current Directory
	file, err := os.Create("./myFile.txt")

	// panic stops the execution and reports the error
	if err != nil {
		panic(err)
	}

	// io.WriteString(fileName , content) returns the length of file
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length of File is ", length)

	// good practice to use DEFER before closing a file
	defer file.Close()

	readingFile("./myFile.txt")

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readingFile(filename string) {

	readBytes, err := ioutil.ReadFile(filename)
	checkNilError(err)

	// fmt.Println("File is being Read \n", readBytes)  // gives ASCII codes as output

	// wrapping the bytes inside a string
	fmt.Println("File is being Read \n", string(readBytes))

}
