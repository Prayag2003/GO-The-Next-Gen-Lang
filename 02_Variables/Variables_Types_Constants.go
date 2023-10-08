package main

import "fmt"

// walrusOperator := 100  #Error
// var walrusOperator int = 100;  # Valid

const PublicVarStartsWithCapitals int = 100 // Public var starts with capital letters

func main() {

	fmt.Println("Variables")
	var userName string = "Gooooo"
	fmt.Println(userName)
	// % T is a placeholder for Type
	fmt.Printf("Variable is of type : %T .\n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T .\n", isLoggedIn)

	var smallInt uint = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type : %T .\n", smallInt)

	var smallFloat float64 = 255.85603287508267
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T .\n", smallFloat)

	// 	default values and some aliases
	var tempVar int
	fmt.Println(tempVar)
	fmt.Printf("Variable is of type : %T .\n", tempVar)

	// Implicit Type
	//  Lexer detects the type of variable
	var name = "Learning-Go"
	fmt.Println(name)

	// No Var Style
	// Using Walrus Operator  ( Var Declaration + Assignment )
	// this operator can be used only inside a Function
	noVarStyleVariable := 100
	fmt.Println(noVarStyleVariable)

	fmt.Println(PublicVarStartsWithCapitals)

}
