package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang ... ")

	languages := make(map[string]string)
	languages["Py"] = "Python"
	languages["Js"] = "JavaScript"
	languages["Cpp"] = "C_Plus_Plus"
	languages["Rb"] = "Ruby"

	/* By Default Sorted and they are not Comma Seperated Values */
	fmt.Println("List of Languages include : ", languages)

	/* Delete a value from the map using Delete Keyword */
	delete(languages, "Rb")
	fmt.Println("List of Languages include : ", languages)

	/* For Each Loop */
	for key, value := range languages {
		fmt.Printf("Key is %v and Value is %v .\n ", key, value)
	}
}
