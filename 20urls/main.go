package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

// go mod init urls
func main() {
	fmt.Println("Welcome to handling URLs in golang.")
	fmt.Println(myurl)

	//parsing the url - extract info
	// result, err := url.Parse(myurl)
	result, _ := url.Parse(myurl) //url module, Parse() method returns result(URL object) and error

	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //lco.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.Port())   //3000
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=ghbj456ghb  - params/reqparams

	qparams := result.Query()                                 //Query() stores the query params in better format
	fmt.Printf("The type of query params are: %T\n", qparams) // The type of query params are: url.Values
	//url.Values - are key value pairs

	fmt.Println(qparams["coursename"]) //[reactjs]
	//since key value pairs , can iterate - order is not guaranteed in this datatype
	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//constructing url from parts of info
	partsOfUrl := &url.URL{ // always pass reference of URL not copy
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String() // another way - string(partsOfUrl)
	fmt.Println(anotherURL)           //https://lco.dev/tutcss
}
