package main

import (
	"fmt"
	"time"
)

func main() {
	go Greetings("Inside GREETING Routine 🤩")
	fmt.Println("\nInside MAIN function 😎")
	time.Sleep(time.Second)
}

func Greetings(s string) {
	fmt.Println("PRAYAG says : ", s)
	fmt.Println("")
	fmt.Println("ByeBye!! Keep learning.")
}
