package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("_____ Files in Go _____")

	createAndWriteFile()
	readFile("./file.txt")
}

// createAndWriteFile creates a new file and writes content to it
func createAndWriteFile() {
	content := "Hello, World! from files"

	// Create a file
	file, err := os.Create("./file.txt")
	checkNilError(err)

	// Write to the file
	length, err := file.WriteString(content)
	checkNilError(err)

	fmt.Printf("File created with %v characters\n", length)

	// Ensure file is closed at the end
	defer file.Close()
}

// readFile reads the content of a given file and prints it
func readFile(filename string) {
	buffer, err := os.ReadFile(filename)
	checkNilError(err)

	fmt.Println("File Content:")
	fmt.Println(string(buffer))
}

// checkNilError checks if an error is not nil and panics if it is
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
Notes:
=========
- **File Handling in Go**
  - `os.Create(filename)`: Creates a file (overwrites if already exists).
  - `file.WriteString(content)`: Writes a string to a file.
  - `os.ReadFile(filename)`: Reads the entire file into memory.
  - `defer file.Close()`: Ensures the file is closed after operations.

- **Error Handling**
  - Always check for errors when working with files to prevent crashes.
  - `panic(err)`: Stops execution if an error occurs (useful for debugging).
  - Use `defer file.Close()` to avoid file resource leaks.

Key Takeaways:
- Always handle errors when working with files.
- Use `defer file.Close()` to ensure files are properly closed.
- `os.ReadFile` is useful for simple file reading.
*/
