package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// TODO: add a mechanism to add middlewares
	// then add a Handle call to add the handler that will call the middlewares in a chain

	http.HandleFunc("/", defaultController)
	http.HandleFunc("/api/", apiController)

	fmt.Println("Ready to listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
