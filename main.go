package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aykay76/ksa/internal/controllers"
)

func main() {
	// TODO: add a mechanism to add middlewares
	// then add a Handle call to add the handler that will call the middlewares in a chain

	http.HandleFunc("/", controllers.DefaultController)
	http.HandleFunc("/api/", controllers.ApiController)

	fmt.Println("Ready to listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
