package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func mainController(w http.ResponseWriter, r *http.Request) {
	// generic controller will load content and do general parsing before handing off to path routed specific controllers
	// TODO: make my root directory configurable
	filename := "./content" + r.URL.Path

	// add default document
	if filename == "./content/" {
		filename = "./content/index.html"
	}

	// TODO: maybe improve the logging a bit ;)
	fmt.Println(filename)

	body, _ := ioutil.ReadFile(filename)

	if strings.HasSuffix(filename, ".css") {
		w.Header().Set("Content-Type", "text/css")
	} else if strings.HasSuffix(filename, ".svg") {
		w.Header().Set("Content-Type", "image/svg+xml")
	} else if strings.HasSuffix(filename, ".html") {
		w.Header().Set("Content-Type", "text/html")

		// convert to string and do some basic SSI
		bodyString := string(body)

		idx := strings.Index(bodyString, "<!--#include file=")
		for idx != -1 {
			idx2 := strings.Index(bodyString, "-->")
			subfile := bodyString[idx+19 : idx2-1]

			subfileContent, _ := ioutil.ReadFile("./content" + subfile)

			newBodyString := bodyString[0:idx] + string(subfileContent) + bodyString[idx2+3:len(bodyString)]
			bodyString = newBodyString
			idx = strings.Index(bodyString, "<!--#include file=")
		}

		body = []byte(bodyString)
	} else if strings.HasSuffix(filename, ".js") {
		w.Header().Set("Content-Type", "text/javascript")
	}

	w.Write(body)
}

func main() {
	http.HandleFunc("/", mainController)

	fmt.Println("Ready to listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
