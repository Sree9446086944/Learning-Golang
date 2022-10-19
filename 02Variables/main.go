package main

import "fmt"

//go mod init variables -> use in production, without module initialization also it works using go run main.go
// creates go.mod file

// jwtToken := 30000 - error since := operator cannot use outside method - accessible by any file in app

//capital letter will make this public variable like public in java
const LoginToken string = "hjhk" // declaring const - cannot change value

func main() {
	//lexer - checks for code if semicolon is there or any other things for validation
	//since intellisence activated, lexer automatically validates
	var userName string = "Sree"
	fmt.Println(userName)
	fmt.Printf("Variable is of type %T .\n", userName) // place holder %T - for type

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type %T .\n", isLogged)

	//https://go.dev/doc/go1.17_spec#Numeric_types

	var smallVal uint8 = 255 //uint8 (byte) - 0 to 255 - unsigned
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T .\n", smallVal)

	var smallFloat float32 = 255.455445112544518855
	fmt.Println(smallFloat) //255.45544   (5 values after decimal) - 32bit values
	fmt.Printf("Variable is of type %T .\n", smallFloat)

	var smallFloats float64 = 255.455445112544518855
	fmt.Println(smallFloats) //255.4554451125445   (more precise) - 64bit values
	fmt.Printf("Variable is of type %T .\n", smallFloats)

	//default values and some aliases
	var anotherVariable int      //just declared not initialized
	fmt.Println(anotherVariable) // 0 - default value if not initialized for int
	fmt.Printf("Variable is of type %T .\n", anotherVariable)

	//implicit type
	var website = "learn.in" //lexer automatically decides what type based on value if not given type, later cannot change value with different type
	fmt.Println(website)
	// website = 3 - error, already lexer assigned string type to website

	//no var style - initialization
	//can use this operator only inside method, not outside/global
	numberOfUser := 300000 // := walrus operator
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T .\n", LoginToken)
}
