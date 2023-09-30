package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Construction of URL !")

	// Type of the URL
	fmt.Printf("Type of URL is %T\n", myURL)
	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

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
