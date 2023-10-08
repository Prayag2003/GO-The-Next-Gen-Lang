package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("MOD in Golang")
	greeter()
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	// running a server
	log.Fatal(http.ListenAndServe(":3000", router))
}

func greeter() {
	fmt.Println("Hello everyone :)")
}

func serveHome(writer http.ResponseWriter, req *http.Request) {
	// to send response
	writer.Write([]byte("<h1>Welcome to Home Page</h1>"))
}
