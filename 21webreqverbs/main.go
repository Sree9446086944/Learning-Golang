package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// go mod init webreqverbs

func main() {
	fmt.Println("Welcome to  web verb video - LCO")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://reqres.in/api/users?page=2"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //close the request

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder //Builder in strings package, used to build string using write() method, efficient, minimize memory copying
	content, _ := ioutil.ReadAll(response.Body)
	//using string builder
	byteCount, _ := responseString.Write(content) //write() takes []byte, now the responseString have content as string
	//Write() returns byte count (content length) and err

	fmt.Println(string(content)) //[]byte to string
	//OR
	fmt.Println("Byte count: ", byteCount)
	fmt.Println(responseString.String()) //prints content

}

// send data to webserver as body
func PerformPostJsonRequest() {
	const myurl = "https://www.thunderclient.com/welcome"
	// make sure sending req to correct url and inform what kind of data sending to them

	// fake json payload
	responseBody := strings.NewReader(`   // can create any kind of data using NewReader(), inside backtick give data
	 {
		"email": "eve.holt@reqres.in",
		"password": "pistol"
	 }
	`)

	//send req
	// http.Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
	response, err := http.Post(myurl, "application/json", responseBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // every req of response is closed

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(string(content))
}

// form data
func PerformFormRequest() {
	const myurl = "https://www.thunderclient.com/welcome"

	//create form data

	data := url.Values{} //type Values map[string][]string - key-value pair
	data.Add("firstname", "Sreekanth")
	data.Add("lastname", "Kumar")
	data.Add("email", "sree@go.com")

	//post req which is url encoded
	//http.PostForm(url string, data url.Values) (resp *http.Response, err error)
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // everything is done connection is closed

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
