package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	const port = 3000
	performGetRequest(port)
	performPostRequest(port)
	performPostFormRequest(port)
}

func performGetRequest(port int) {
	fmt.Println("\n\nYou are inside get Request")
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
	fmt.Println("\n\nYou are inside Post Request")
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

func performPostFormRequest(port int) {
	fmt.Println("\n\nYou are inside Post Form Request")
	myurl := fmt.Sprintf("http://localhost:%d/postform", port)

	// creating fake formdata
	data := url.Values{}

	data.Add("age", "20")
	data.Add("firstname", "Prayag")
	data.Add("lastname", "Bhatt")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	respBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(respBody))

}
