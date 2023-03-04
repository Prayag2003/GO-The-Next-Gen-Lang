package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Construction of URL !")

	// Type of the URL
	fmt.Printf("Type od URL is %T", myURL)

	// Since the URL is nothing but Key-Value Pairs , we can construct one via

	// always use & before url ***

	newURL := &url.URL{
		Scheme:  "http",
		Host:    "prayagbhatt.in",
		Path:    "learn",
		RawPath: "user=Prayag",
	}

	// Converting it into a String
	URLString := newURL.String()
	fmt.Println(URLString)

}
