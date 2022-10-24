package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// go mod init mymath

func main() {

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("The sum is: ", myNumberOne+int(myNumberTwo)) //automatically cast to int, by the float64 loose its precision

	//random number - by Math and by cryptography algorithm
	// https://pkg.go.dev/math/rand#Intn
	// https://pkg.go.dev/crypto/rand#Int

	// rand.Seed(7)              //rand algorithm is driven by seed, pass any value so will generate random number, if not given always 1
	//or pass time as value
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) //number exclusive, i.e 0 to 4 range, so +1 -> 1 to 5

	//random from crypto - crypto/rand

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5)) // Int(reader,bigInt)  -> 0 to 4
	fmt.Println(myRandomNum)

}
