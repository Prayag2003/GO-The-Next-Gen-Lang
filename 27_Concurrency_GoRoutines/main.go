package main

import (
	"fmt"
	"time"
)

func main() {
	go HelloWorld("Hello")
	HelloWorld("World")
}

func HelloWorld(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Prayag says : ", s)
	}
	fmt.Println("What")
}
