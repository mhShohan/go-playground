package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("_____files_________")

	createAndWriteFile()
	readFile("./file.txt")
}

func createAndWriteFile() {

	content := "Hello, World! from files"

	// Create a file
	file, err := os.Create("./file.txt")

	checkNillError(err)

	length, err := file.WriteString(content)

	checkNillError(err)

	fmt.Printf("File created with %v characters\n", length)
	defer file.Close()
}

func readFile(filename string) {
	buffer, err := os.ReadFile(filename)

	checkNillError(err)

	fmt.Println(string(buffer))

}

// checkNillError checks if the error is not nil and panics
func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}
