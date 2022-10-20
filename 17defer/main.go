package main

import "fmt"

// go mod init defer
func main() {

	//defer
	// programs execute line by line, when any statements is given defer keyword, that line is cut and it will execute at end of fn in last in firstout way

	// defer fmt.Println("World") // skips and this line execute after all other statements
	// fmt.Println("Hello")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	//Hello , Two, One, world  -> executes non defered statements from flow first, then defered statement execute lifo (last in first out)
	//all defered values get stored in stack, and at end execute LIFO
	myDefer()
	//flow
	//    1st stack - "World","One","Two"
	//    non defer - "Hello" //executes
	//    2nd stack - 0,1,2,3,4
	//result -> Hello , 4,3,2,1,0,Two,One,World   -> LIFO(Hello, 2ndstack Lifo, 1ststack lifo)
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
