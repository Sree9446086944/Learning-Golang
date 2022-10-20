package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// go mod init files
func main() {
	fmt.Println("Files in golang")

	content := "This needs to go in file."
	file, err := os.Create("./mylcogofile.txt") // os package - Create() method creates file, returns file and error, ./- current directory

	// if err != nil {
	// 	panic(err) //panic() shut down program and shows err given inside fn
	// }
	//or
	checkNilError(err)
	length, err := io.WriteString(file, content) // io package - WriteString() method writes string in file, returns length and error

	fmt.Println("Length is: ", length)
	// file.Close() //close file
	defer file.Close() //recommended way, since execute LIFO if multiple files , at the end

	//read file
	readFile("./mylcogofile.txt") //in current directory(./)
}

func readFile(fileName string) {
	//whenever read files, it not read into string, it always read into bytes
	databyte, err := ioutil.ReadFile(fileName)

	// if err != nil {
	// 	panic(err)
	// }
	//or
	checkNilError(err)
	fmt.Println("Text data inside file is \n", databyte) ///[84 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105 110 32 102 105 108 101 46]

	// to handle bytes and convert to string
	fmt.Println("Text data inside file is \n", string(databyte)) //Text data inside file is This needs to go in file.
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
