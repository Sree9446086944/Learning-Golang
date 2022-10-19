package main

import (
	"fmt"
	"sort"
)

// go mod init slices
func main() {
	//slices
	//under the hood arrays, a layer of abstraction with features inside - slices

	// var fruitList = []string{}                         //no size define, also need to initialize {}
	// fmt.Printf("Type of fruitlist is %T\n", fruitList) //[]string - means slices

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana") //append(sliceToAppend, elements ...Type) - to add elements to slices

	fmt.Println(fruitList) //[Apple Tomato Peach Mango Banana]

	fruitList = append(fruitList[1:]) // [Tomato Peach Mango Banana] - slices the slice
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)            //[Peach Mango]
	fruitList = append(fruitList[:3]) //[Peach Mango Banana]
	fmt.Println(fruitList)

	//make keyword for allocating slice
	highScores := make([]int, 4) //make(type,size) - declare and initializes
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	fmt.Println(highScores) //[234 945 465 867]
	// highScores[4] = 777 //error - out of bound

	// given slice and size 4, this is default allocation  till index 3, when using append,
	//  this will reallocate memory and all values get accomadated,
	// memory allocation happens again
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores) ///[234 945 465 867 555 666 321]

	//sort package - used for slices not array
	fmt.Println(sort.IntsAreSorted(highScores)) // checks slice is sorted - false
	sort.Ints(highScores)                       //sorts int slices , many operations in sort pkg
	fmt.Println(highScores)                     //[234 321 465 555 666 867 945]
	fmt.Println(sort.IntsAreSorted(highScores)) //true

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2 //index 2 remove
	//append also used for removing element
	courses = append(courses[:index], courses[index+1:]...) //0 to 1, 3 to end
	fmt.Println(courses)                                    //[reactjs javascript python ruby]
}
