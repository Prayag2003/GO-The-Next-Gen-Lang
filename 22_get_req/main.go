package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	const port = 3000
	performGetRequest(port)
	performPostRequest(port)
}

func performGetRequest(port int) {
	url := fmt.Sprintf("http://localhost:%d/get", port)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("The Status code is ", response.StatusCode)
	fmt.Println("The Content length is ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(content)
	fmt.Println(responseString.String())
	fmt.Println("Bytecount is ", byteCount)

	// fmt.Println(string(content))
}

func performPostRequest(port int) {
	url := fmt.Sprintf("http://localhost:%d/post", port)

	// TODO: fake json payload

	requestBody := strings.NewReader(`
		{
			"name":"prayag",
			"hobby":"play chess, piano, listen to songs, draw sketches/paintings"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
