package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:3000/get"
	fmt.Println("GET Request in Go ... ")
	PerformGetRequest(&url)

}

func PerformGetRequest(myURL *string) {

	// Making a Web request
	response, err := http.Get(*myURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // it's user's responsibility to close the response once the task is done

	fmt.Println("Status Code is ", response.StatusCode)
	fmt.Println("Content Length is ", response.ContentLength)

	// reading all the content from the body of the Response
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	// wrapping the content into a string so as to avoid displaying of the bytes
	fmt.Println(string(content))

	/* Another Method to Write the Content into String using Strings.builder */

	// https://pkg.go.dev/strings#Builder

	var responseString strings.Builder
	bytecount, _ := responseString.Write(content)
	fmt.Println("No of Bytes in the String is ", bytecount)

	printContent := responseString.String()
	fmt.Println("Content is ", printContent)

}
