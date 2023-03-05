package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post Request in GO ...")

	// Data is sent / posted to the BackEnd mostly in two forms :-
	// 1) URL- Encoded Form
	// 2) JSON Format

	myURL := "http://localhost:3000/post"
	RequestPostJSON(&myURL)

}

func RequestPostJSON(myURL *string) {

	// fake JSON payload
	// we can create it using NewReader Methods of the strings library

	payload := strings.NewReader(`

	{
		"courseName" : "Learning GO ",
		"Price" : "0" ,
		"Platform" : "www.learncodeonline.com"
	}
`)

	/************  POSTING  **************/

	// In the post request , one needs to focus on 3 things
	// 1) URL
	// 2) Content-Type ( can be found in Header section in the Thunder Client)
	// 3) Request Body

	// http.Post( URL  , Content-Type , Request Body)

	response, err := http.Post(*myURL, "application/json", payload)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	readRequest, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(readRequest))

}
