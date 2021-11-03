package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// make this a middleware component that prepares the content header etc, then pass to additional middleware
// that will do data prep etc.
func DefaultController(w http.ResponseWriter, r *http.Request) {
	contentPath := "./web"

	filename := contentPath + r.URL.Path

	// add default document
	if filename == contentPath+"/" {
		filename = contentPath + "/index.html"
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
			right := bodyString[idx:]
			idx2 := strings.Index(right, "-->") + idx
			subfile := bodyString[idx+19 : idx2-1]

			subfileContent, _ := ioutil.ReadFile(contentPath + subfile)

			newBodyString := bodyString[0:idx] + string(subfileContent) + bodyString[idx2+3:]
			bodyString = newBodyString

			idx = strings.Index(bodyString, "<!--#include file=")
		}

		body = []byte(bodyString)
	} else if strings.HasSuffix(filename, ".js") {
		w.Header().Set("Content-Type", "text/javascript")
	}

	w.Write(body)
}
