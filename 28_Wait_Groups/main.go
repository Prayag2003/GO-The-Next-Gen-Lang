package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually these are pointers
var mut sync.Mutex

var urls = []string{"https://github.com/Prayag2003"}

func main() {
	websiteList := []string{
		"https://youtube.com",
		"https://www.google.com/",
		"https://go.dev",
	}
	for _, weblist := range websiteList {
		go getStatusCode(weblist) // no result since we are not waiting the thread the execute and finish its task before closing the main method
		wg.Add(1)                 // adding a counter of number of routines in work
	}

	wg.Wait() // blocks the main method, waits for routines to finish
	fmt.Println(urls)
}

func getStatusCode(endpoint string) {

	defer wg.Done() // signal that the routine has completed it's own task
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Endpoint is incorrect")
		panic(err)
	}

	// the write operation requires a lock on it until it finishes writing into memory
	mut.Lock()
	urls = append(urls, endpoint)
	mut.Unlock()

	fmt.Printf("Status Code : %d for %s\n", res.StatusCode, endpoint)
}
