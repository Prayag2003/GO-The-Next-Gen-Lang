package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func main() {
	var wg sync.WaitGroup

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := []int{}

	for _, data := range data {
		wg.Add(1)
		go operation(&wg, &res, data)
	}
	wg.Wait()
	fmt.Println(res)
}

func operation(wg *sync.WaitGroup, res *[]int, data int) {

	lock.Lock()

	// NOTE: Critical Section begins
	defer wg.Done()
	processedData := data + 10
	*res = append(*res, processedData)
	// NOTE: Critical Section ends

	lock.Unlock()

}
