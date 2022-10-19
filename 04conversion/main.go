package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza b/w 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for ratiing, ", input)

	//to do maths operations, convert input string to int/float
	// strconv.ParseFloat(input,64) //- input string have \n at end, so need to trip spaces to get value and parse else error eg :- "5\n"
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //ParseFloat takes string and bitsize (32/64) ,methods capital to make it public

	if err != nil { // if err has a value
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to rating: ", numRating+1)
	}

}
