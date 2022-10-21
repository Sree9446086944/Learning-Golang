package main

import (
	"encoding/json"
	"fmt"
)

// go mod init json

func main() {
	fmt.Println("Welcome to json.")
	//encoding of json - datas like arrays/slices/keyValues etc convert to a json
	EncodeJson()
}

// not exporting/public so lowercase
type course struct {
	Name     string `json:"coursename"` //alias as 3rd param, provided by json package, when convert to json, this alias is seen as key name
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // - means dont need this fields to get reflected, will not be shown in json
	Tags     []string `json:"tags,omitempty"` //omitempyty means , if field is null/nil , dont show the field
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 299, "sree@go.com", "abc", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "sree@go.com", "bcd", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "sree@go.com", "sre", nil},
	}

	//package this data as JSON data

	finalJson, err := json.Marshal(lcoCourses) //json package Marchal() returns JSON encoding of object
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson) // array of struct objects in json

	//MarshalIndent()
	// json.MarshalIndent(v any, prefix string, indent string) ([]byte, error)
	//internally when all data converted into json, it is indented based on \t(tab)
	//if prefix given, that will be shown in every line start
	finalJson2, err := json.MarshalIndent(lcoCourses, "", "\t") // formatted
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson2) // formatted json

}
