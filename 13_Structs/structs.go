package main

import "fmt"

func main() {
	fmt.Println("Structs in  Golang ...")

	/* There is no INHERITANCE or SUPER or PARENT */
	Prayag := User{"Prayag Bhatt", "prayagbhatt2003@gmail.com", "123456789", true}
	fmt.Println("Details of Prayag includes : ", Prayag)

	/* Better Way of Priting */
	fmt.Printf("User Details are %+v \n: ", Prayag)
	fmt.Printf("User Name is %+v \n : ", Prayag.Name)

}

/* Struct Name as well as it's property names are capital since they will used Publically" */

type User struct {
	Name     string
	Email    string
	Phone_No string
	Status   bool
}
