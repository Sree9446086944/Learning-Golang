package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// go mod init webrequests

const url string = "https://lco.dev"

func main() {
	//grab data from webpage
	//http package - https://pkg.go.dev/net/http

	//https://pkg.go.dev/net/http#Response - type response - when request is made neither we read/write closing that request
	//even if we change pade or do anything, it still does since not closing connection, so we need to close that response
	//Response struct have Close property

	fmt.Println("LCO web requuests.")
	//http requests
	response, err := http.Get(url) //http from net package, err if any n/w errors

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response) //Response is of type: *http.Response (Reference of original response we get - pointer)
	//guarantee that not copy of response is got, actually getting response so that we can manipulate - *http.Response

	defer response.Body.Close() // callers responsibility to close the connection, close when everything done - defer

	//read the response - majority of read is done by ioutil
	databytes, err := ioutil.ReadAll(response.Body) // there is Body, headers attributes etc
	if err != nil {
		panic(err)
	}
	// fmt.Println(databytes)  //byte slice
	content := string(databytes) //contents of web page as string, entire html
	fmt.Println(content)
}
