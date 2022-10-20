package main

import "fmt"

// go mod init methods
func main() {

	//functions inside class/structs in go - methods
	sree := User{"Sree", "sree@go.com", true, 25}
	//calling methods of struct
	sree.GetStatus() //Is user active:  true

	fmt.Printf("Mail is: %v\n", sree.Email) //Mail is: sree@go.com
	//mail change using method
	sree.NewMail()                          //Email of this user is:  test@go.com
	fmt.Printf("Mail is: %v\n", sree.Email) //Mail is: sree@go.com
	// mail didnt change i.e, copy of objects/structs is passed on method, so to pass original object,
	// need to pass reference of that object or pointer
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//creating methods
//methods should be part of struct, need to pass Struct or property of struct
//capital to export/public
func (u User) GetStatus() { //func (struct type/struct property) nameOfStruct
	fmt.Println("Is user active: ", u.Status)
}

//similar to setter
func (u User) NewMail() {
	u.Email = "test@go.com"
	fmt.Println("Email of this user is: ", u.Email)
}
