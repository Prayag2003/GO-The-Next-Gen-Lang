package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Prayag2003/connecting_go_and_mongo/router"
)

func main() {

	fmt.Println("Connecting DB and Creating API for a Cinema")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening at port 3000")
}
