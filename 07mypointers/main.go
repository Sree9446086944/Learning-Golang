package main

import "fmt"

// go mod init mypointers
func main() {

	//sometimes when we pass variables, copy of the var is passed.whenever there is case to avaoid this and
	// guarantee that always actual value should be passed on, then prefer pointer to passed on(i.e memory addess is passed which guarantee that actual value is passed)

	//pointer - direct reference to memory address
	//since we pass memory address(while passing pointer) ,guarantees actual value is passed

	//declare pointer
	// var ptr *int                             // * tell this datatype is of pointer, *int pointer type responsible for holding integers
	// fmt.Println("Value of pointer is ", ptr) // nil , if not given value to pointer, value is nil, pointer pointing to nothing, no memory address

	myNumber := 23 // auto infer its int

	var ptr = &myNumber                              //creating pointer which referencing to memory address of myNumber , reference(&)
	fmt.Println("Value of actual pointer is ", ptr)  //0xc00001c0a8 - actual memory address of myNumber
	fmt.Println("Value of actual pointer is ", *ptr) // 23 - actual value inside it

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber) //New value is:  25
	fmt.Println(*ptr)                       //25

	//pointer makes it guarantee that whatever operations with value, those operations are performed on actual value not on copy of value
}
