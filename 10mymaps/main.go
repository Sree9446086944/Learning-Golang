package main

import "fmt"

// go mod init mymaps
func main() {
	//maps
	// key-value format, retrival is easy since we know what is exactly to fetch

	languages := make(map[string]string) //create map using map which also initializes, map[keyType]valueType

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages) //map[JS:Javascript PY:Python RB:Ruby]
	fmt.Println("JS shorts for: ", languages["JS"])   //JS shorts for:  Javascript

	//delete values
	delete(languages, "RB")                           //delete keyword is used to remove values, can also used for slices, here RB key/value get removed
	fmt.Println("List of all languages: ", languages) //map[JS:Javascript PY:Python]

	//loops for map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value) //%v placeholder - for value itself
	}
}
