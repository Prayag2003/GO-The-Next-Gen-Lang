package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "http://localhost:8000"
	fmt.Println("GET Request in Go ... ")
	PerformGetRequest(&url)

}

func PerformGetRequest(myURL *string) {

	response, err := http.Get(*myURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Generated Response is ", response.StatusCode)
	fmt.Println("Content Length is ", response.ContentLength)

}
