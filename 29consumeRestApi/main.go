package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// http req to remote server
func main() {
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("The http request failed with error %s\n", err)
	} else {
		//read from response
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	//post req
	//create req body
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	//marshal the map to json
	jsonValue, _ := json.Marshal(jsonData)

	//tell what kind of data we r sending using Content-Type
	//:= not used since response and err already defined above
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue)) //buffer json into bytes and send req body
	if err != nil {
		fmt.Printf("The http request failed with error %s\n", err)
	} else {
		//read from response
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	//post req  - another way
	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}             //Client is a struct
	response1, err := client.Do(request) //send req with client
	if err != nil {
		fmt.Printf("The http request failed with error %s\n", err)
	} else {
		//read from response
		data, _ := ioutil.ReadAll(response1.Body)
		fmt.Println(string(data))
	}
}
