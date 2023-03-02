package main

import "fmt"

func main() {

	var input, input2 int
	fmt.Println("Functions In Loop ... ")
	Greet()
	Greet2()

	fmt.Println("Enter 2 Numbers :")
	fmt.Scanf("%v\n%v ", &input, &input2)
	res := adder(input, input2)
	fmt.Printf("Sum Of %v and %v is %v.\n", input, input2, res)

	/* Multiple Return values */
	ans, MyMessage := multiAdder(10, 20, 30, 40)
	fmt.Println("Multi Adder sums as ", ans)
	fmt.Println("My Message is ", MyMessage)

}

func Greet() {
	fmt.Println("Greetings !! ")
}

func Greet2() {
	fmt.Println("Greetings 2.0 !! ")
}

func adder(one int, second int) int {
	return one + second
}

func multiAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hello this is sum\n"
}
