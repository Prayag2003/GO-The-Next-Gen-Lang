package main

import (
	"fmt"
	"sort"
)

func main() {

	// Slices :- When Arrays get Abstraction
	fmt.Println(" Slices in Golang ... ")

	// initialising a Slice
	var fruitList = []string{}
	fmt.Printf("Type of FruitList is : %T\n ", fruitList)

	fruitList = append(fruitList, "Apple", "Mango", "Banana", "Peach")
	fmt.Println("FruitList includes : ", fruitList)

	/* Slicing the Slice */
	//  [ inclusive starting : exclusive end ]
	//  changes the original slice
	fruitList = append(fruitList[2:])
	fmt.Println(fruitList)

	// M A K E  :- allocates fixed size of memory
	marksOfSubjects := make([]int, 4)
	marksOfSubjects[0] = 100
	marksOfSubjects[1] = 92
	marksOfSubjects[2] = 88
	marksOfSubjects[3] = 79
	// marksOfSubjects[4] = 99
	// fmt.Println(marksOfSubjects) // Out of Range Error

	// However using Append won't give any errors since append reallocates new memory
	marksOfSubjects = append(marksOfSubjects, 70, 90, 80)
	fmt.Println(marksOfSubjects)

	// Sorting
	sort.Ints(marksOfSubjects)
	fmt.Println(marksOfSubjects)

	// IsSorted
	ans := sort.IntsAreSorted(marksOfSubjects)
	fmt.Printf("Is the Slice Sorted ??? ")
	if ans == false {
		fmt.Println("Nope")
	} else {
		fmt.Println("Yes")
	}

	// ****** Removing a Certain Value *******
	var courses = []string{"C++", "JavaScript", "Golang", "Java", "Python", "Rust", "C#", "NodeJS", "Solidity"}
	fmt.Println(courses)

	// Suppose we need to remove C# from the list , it has index of 6
	index := 6

	// we will using append slicing on the courses slice , we remove the index element and using ... for multiple arguments
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
