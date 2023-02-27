package main

import "fmt"

func main() {

	fmt.Println("Arrays in Golang ... ")

	// Array length has to be specified at the time of creation , else err
	var subjects [5]string
	subjects[0] = "Design Analysis And Algos  "
	subjects[1] = "Machine Learning  "
	subjects[2] = "Cybersecurity  "
	subjects[4] = "Operating Systems  "

	fmt.Println("Subject List includes : ", subjects)
	fmt.Println("Length of Subject List is : ", len(subjects))

	var friends = [2]string{"abc", "pqr"}
	fmt.Println("Subject List includes : ", friends)

}
