package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {

	fmt.Println("Handling Web Requests in GO !!! ")

	/*******************  H   T   T   P  ************************/

	/*  Client (Browsers like Brave , Chrome ) sends HTTP request to the website
	( Server ) , Server recieves the request and processes it out in the application layer

	Server sends back HTTP Responsae in Form of objects ...
	*/

	/* Neither ReadResponse nor the Response.Write ever closes a connection
	... It has to be closed by Us ... */

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The response is of type : %T\n", response)
	fmt.Println("The response is ", response)

	// closing the connection
	defer response.Body.Close()

	// reading the Response , majorly by the IOUTIL
	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// Wrapping Databytes into a string so as we can get in a readable form instead of ASCII Values
	responseInString := string(databytes)
	fmt.Println("\n\nResponse is : ", responseInString)

}
