package main

import "fmt"

type NumType interface {
	int16 | int32 | int64 | float32 | float64
}

func sumSlicesWOGenerics(slice1 []int16) int16 {
	var ans int16
	for i := range slice1 {
		ans += slice1[i]
	}
	return ans
}

func sumSliceWithGenerics[T NumType](nums []T) T {
	var ans T
	for i := range nums {
		ans += nums[i]
	}
	return ans
}

func main() {

	slice1 := []int16{1, 2, 3, 4, 5}
	slice2 := []int32{1000, 200, 300, 400, 500}
	slice3 := []int64{1e6, 1e9, 1e10}
	slice4 := []float32{1.234, 3.4565, 4220.464}
	slice5 := []float64{100.111111, 200.222222, 300.333333}

	fmt.Println("\nWithout Using Go Generics")
	fmt.Println(sumSlicesWOGenerics(slice1))

	fmt.Println("\nUsing Go Generics")
	fmt.Println(sumSliceWithGenerics(slice1))
	fmt.Println(sumSliceWithGenerics(slice2))
	fmt.Println(sumSliceWithGenerics(slice3))
	fmt.Println(sumSliceWithGenerics(slice4))
	fmt.Println(sumSliceWithGenerics(slice5))

}
