package main

import (
	"fmt"
	"math/rand"
	"time"
)

// go mod init switchcase
func main() {
	fmt.Println("Switch and case in golang.")

	rand.Seed(time.Now().UnixNano()) // generate random number every time precise
	diceNumber := rand.Intn(6) + 1   // b/w 0 to 5,  add 1 to get b/w 1 to 6

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open.")
	case 2:
		fmt.Println("You can move 2 spot.")
		// fallthrough // this continue to print next case also
	case 3:
		fmt.Println("You can move 3 spot.")
	case 4:
		fmt.Println("You can move 4 spot.")
	case 5:
		fmt.Println("You can move 5 spot.")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again.")
	default:
		fmt.Println("What was that!")
	}
}
