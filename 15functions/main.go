package main

import "fmt"

// go mod init functions
func main() {
	// main() is the entry point, we dont need to call the fn
	//if u declare a function, need to call the fn to execute
	//cannot define fn inside a fn
	fmt.Println("Welcome to functions in golang.")
	greeter()

	//signatures - what kind of value passed and what kind of value returned back
	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	// proRes := proAdder(2, 3, 4, 5)
	// proRes, _ := proAdder(2, 3, 4, 5)  //  if dont need value _
	proRes, myMessage := proAdder(2, 3, 4, 5)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Message is: ", myMessage)
}

//when we dont know how many values passed - varargs
//values is slice now
//... -> called veriadic functions
func proAdder(values ...int) (int, string) { //can return any number of values, place inside ()
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi"
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Hi")
}
