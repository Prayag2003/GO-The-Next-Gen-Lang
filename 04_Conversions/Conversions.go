package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate our Pizzas Between 1 to 5 : ")
	input, _ := reader.ReadString('\n')

	fmt.Println(" User Rated : ", input)
	fmt.Printf(" Before Conversion ,  Rating Format is : %T  \n", input)

	// Converting String to Int / Float using #TODO: strconv ( String Conversion to some other type )

	// ratingInFloat, err := strconv.ParseFloat(input, 64) // gives an error since Strings also consists of \n at the end

	// we use #TODO: strings package to correct the above error
	// TrimSpace removes the spaces
	ratingInFloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(" Rating is : \n", ratingInFloat)
		fmt.Printf(" After Conversion , Rating Type is %T ", ratingInFloat)
	}
}
