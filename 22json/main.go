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

	//decoding json data - means consume json data
	DecodeJson()
}

// not exporting/public so lowercase
type course struct {
	Name     string `json:"coursename"` //alias as 3rd param, provided by json package, when convert to json, this alias is seen as key name
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // - means dont need this fields to get reflected, will not be shown in json
	Tags     []string `json:"tags,omitempty"` //omitempyty means , if field is null/nil , dont show the field
}

// encode json - convert to json
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

// decode json - json to object/struct
func DecodeJson() {
	//data we get from web will be []byte (slice of bytes)
	//create fake json
	//used []byte since real web res is in []byte format
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS",
		"Price": 299,
		"website": "sree@go.com",
		"tags": [
				"web-dev",
				"js"
		        ]
    }
	`)

	//data coming from web, create struct for that and put in struct
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb) // checks if []byte is a valid json format
	if checkValid {
		fmt.Println("JSON was valid")
		//lcoCourse interface/struct might be a copy passed, so pass reference to guarantee since we need to store data in lcoCourse
		//  if needed, give & to sure
		// properties of struct get populated
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v", lcoCourse)
	} else {
		fmt.Println("JSON was not valid.") // main.course{Name:"ReactJS", Price:299, Platform:"sree@go.com", Password:"", Tags:[]string{"web-dev", "js"}}

	}

	//some cases where you want to add data to key value (map)

	var myOnlineData map[string]interface{}        //give interface{} as valuetype since dont know what kind of value coming(might be string,int,array etc)
	json.Unmarshal(jsonDataFromWeb, &myOnlineData) //pass reference of myOnlineData, since to add values to original myOnlineData var
	fmt.Printf("%#v\n", myOnlineData)
	// main.course{Name:"ReactJS", Price:299, Platform:"sree@go.com", Password:"", Tags:[]string{"web-dev", "js"}}map[string]interface {}{"Price":299, "coursename":"ReactJS", "tags":[]interface {}{"web-dev", "js"}, "website":"sree@go.com"}

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
