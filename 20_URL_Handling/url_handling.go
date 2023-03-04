package main

import (
	"fmt"
	"net/url"
)

// fictitious URL
const myurl = "http://prayagbhatt:3000/dir/aboutMe?languageexpertise=Cplusplus&others=GO"

func main() {
	fmt.Println("Handing URL....")
	fmt.Printf("Type of the URL is %T\n", myurl)
	fmt.Println(myurl)

	// Parsing ( Meaning :- Analysing into syntatical components)
	// using URL Module
	result, _ := url.Parse(myurl)
	fmt.Println(result)

	// components of the URL
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.OmitHost)

	qparams := result.Query()
	fmt.Println(qparams)

	for _, value := range qparams {
		fmt.Println("Param is ", value)
	}

}
