package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Let's Do Some Arithmentic ...")
	var num1 int = 10
	var num2 int = 4
	var num3 float32 = 100
	var num4 int = 30
	var num5 int = 0

	ans := num1 / num2
	fmt.Println(ans)

	ans2 := num3 / float32(num4) // conversion
	fmt.Println(ans2)

	ans3 := int32(num3) % 10
	fmt.Println(ans3)

	ans4 := math.Sqrt(100)
	fmt.Println(ans4)

	// Run-Time Error
	ans5 := num4 / num5
	fmt.Println(ans5)
	panic(ans5)

}
