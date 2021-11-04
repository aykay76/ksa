package restclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getToken() string {
	body, _ := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	return string(body)
}

func addHeaders(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+getToken())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
}

func getHttpClient() *http.Client {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	certs, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/ca.crt")
	if err != nil {
		log.Fatalf("Failed to append to RootCAs: %v", err)
	}

	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	return client
}

func get(url string) []byte {
	var req *http.Request
	var resp *http.Response
	var err error

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	addHeaders(req)

	httpClient := getHttpClient()

	resp, err = httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	responseBody, _ := ioutil.ReadAll(resp.Body)

	return responseBody
}

func post(url string, body interface{}) []byte {
	var req *http.Request
	var resp *http.Response
	var err error
	var bodyBytes []byte

	bodyBytes, err = json.Marshal(body)
	if err != nil {
		fmt.Println(err)
	}

	req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}

	addHeaders(req)

	httpClient := getHttpClient()

	resp, err = httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	responseBody, _ := ioutil.ReadAll(resp.Body)

	return responseBody
}
