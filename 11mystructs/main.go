package main

import "fmt"

// go mod init mystructs
// no inheritance in go, No super or parent
func main() {

	//use struct
	sree := User{"Sree", "sree@go.com", true, 25}
	fmt.Println(sree)                       // {Sree sree@go.com true 25}
	fmt.Printf("Sree details: %+v\n", sree) // {Name:Sree Email:sree@go.com status:true Age:25} , %+v place holder for formating struct - gives structure

	fmt.Printf("Name os %v and email is %v.\n", sree.Name, sree.Email)

}

//define struct - custom type
//like a class or template, capital User since need to export
type User struct {
	//fields also can be accessed anywhere since capital letter
	Name   string
	Email  string
	status bool
	Age    int
}
