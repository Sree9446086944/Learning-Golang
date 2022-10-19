package main

import "fmt"

// go mod init myarray

func main() {
	var fruitList [4]string // basic declaration of array

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)            // [Apple Tomato  Peach]  space b/w 1 and 3 index since 2nd index not given, default "" empty string
	fmt.Println("Length of fruitList: ", len(fruitList)) //4 always since it is declared size 4 initially

	var vegList = [5]string{"potato", "beans", "mushroom"} //declared and initialized with values
	fmt.Println("Veggie list is: ", vegList)
	fmt.Println("Veggie length: ", len(vegList)) //5
}
