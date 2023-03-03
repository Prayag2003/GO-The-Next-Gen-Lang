package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang ... ")
	Prayag := User{"Prayag", 20, true, "pqr@gmail.com"}
	fmt.Println(Prayag)
	Prayag.GetStatus()
	Prayag.updateMail()
	fmt.Println(Prayag)
}

type User struct {
	Name   string
	Age    int
	Status bool
	email  string
}

// passing struct inside function parameters

// func (struct with its object ) name(){

// }

func (u User) GetStatus() {
	fmt.Println("Is User Online ? ", u.Status)

}

/* Updating the enitity */
// wont reflect any changes in the User Prayag since the struct's Copy is being passed to the method
func (u User) updateMail() {
	u.email = "xyz@gmail.com"
	fmt.Println("New Email is ", u.email)
}
