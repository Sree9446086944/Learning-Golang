package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	// https://pkg.go.dev/bufio  - this is buffer that can read from i/p , o/p - keyboard or other
	//take i/p from std input/ keyboard

	//used := operator since dont know what type comes from i/p
	//NewReader and Stdin is capital to make it public
	reader := bufio.NewReader(os.Stdin) // bufio pkg have NewReader to create reader, to tell from where to read, use os package Stdin(standard i/p - keyboard)
	//println() automatically inject \n at end
	fmt.Println("Enter the rating for our pizza:")

	// comma ok syntax or err ok syntax

	//in  go , no try catch, it treats error as variable or value

	// input,err := reader.ReadString('\n') // read string till newline()\n
	input, _ := reader.ReadString('\n') // if dont care about err
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input) //string

	//when u r reading something from stdinput, there is chace that something might go wrong, error comes in we can store in variable
	// input,err := reader.ReadString('\n')   // comma error syntax
	//if everything right, we get input in 'input' variable, if error comes, catch in err
	//like try and catch in other language
}
