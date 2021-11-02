package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainController)
	http.HandleFunc("/api/", apiController)

	fmt.Println("Ready to listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
