package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mahi160/hiteflix/routers"
)

func main() {
	r := routers.Router()
	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server started, listening on port 4000")
}
