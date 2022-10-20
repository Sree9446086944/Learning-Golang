package main

import "fmt"

// go mod init loops
func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	//type 1
	for d := 0; d < len(days); d++ { // no ++d in go , only d++
		fmt.Println(days[d])
	}

	//type 2
	for i := range days { // i is index
		fmt.Println(days[i])
	}

	//type 3
	// range return index and value
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	//if need value only
	for _, day := range days {
		fmt.Printf(" value is %v\n", day)
	}

	rougueValue := 1
	//similar to while in other languages
	for rougueValue < 10 {

		// if rougueValue == 5 {
		// 	break // entire loop breaks, doesnt allow to continue
		// }

		if rougueValue == 5 {
			rougueValue++
			continue // skips one phase/loop, if value is not incremented then it will come inside this condn again and again
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	// goto statements

lco: // label / keyword, when given : after keyword it points to goto - can give any name for label
	fmt.Println("Jumping")

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco // goto any label
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

}
