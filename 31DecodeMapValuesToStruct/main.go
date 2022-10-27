package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// map values into struct/generic datastructure
func main() {
	//mapstructure - package
	// go get github.com/mitchellh/mapstructure

	mapPerson := make(map[string]interface{}) //-> key - string, value - interface , make() declare and initialize
	mapPerson["firstname"] = "Nic"            // interface with string value
	mapPerson["lastname"] = "Raboy"

	mapAddress := make(map[string]interface{})
	mapAddress["city"] = "Tracy"
	mapAddress["state"] = "CA"
	mapPerson["address"] = mapAddress // interface with map value

	fmt.Println(mapPerson) // map[address:map[city:Tracy state:CA] firstname:Nic lastname:Raboy]

	//convert map to datastructure
	var person Person
	mapstructure.Decode(mapPerson, &person) //decode mapPerson map and store in person struct, &person since original person var is to be changed
	fmt.Println(person)                     //{Nic Raboy {Tracy CA}}
	// mapstructure.DecodeMetadata(mapPerson, &person, &mapstructure.Metadata{})
	// fmt.Println(person)
}

type Person struct {
	Firstname string
	Lastname  string
	Address   struct {
		City  string
		State string
	}
}

/*type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
}*/
