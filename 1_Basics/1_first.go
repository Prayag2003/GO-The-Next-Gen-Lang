// 1 ) Compiled language and has different executables for Different OS

// 2 ) It is an OOPs lang but doesn't contain Classes and Objects , rather has Structs instead classes , also concept of overloading and overriding is not present

/* package main tells us that out code should be compiled into an StandAlone Executable File */

package main

import "fmt"

func main() {
	fmt.Println("Hello , I am Learning GOLANG !!! ")

	// LEXER ensures the rules of Grammar are properly followed
	// Also takes care of semi-colons ;)

	/*   Capital Letters of the Methods indicate that when they are exported , they are public   */
	fmt.Println("Lexers --> Adds Semi-colons ;) ... ")

	/* STACTICALLY DEFINED TYPE : once a variable type is declared , it is fixed and hence type cannot be changed

	Almost Everything in the Golang is a type  */

	/*
		Types Include
		1) String                6) Arrays
		2) Boolean               7) Slices
		3) Integer               8) Maps
		4) Floating              9) Structs
		5) Complex              10) Pointers

		-- Functions , Channels are types as well
	*/
}
