package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter Your Name: ")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput) // Remove newline character

	fmt.Println("Enter Your Age: ")
	ageReader := bufio.NewReader(os.Stdin)
	ageInput, _ := ageReader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput) // Remove newline character

	fmt.Println("Hello, " + nameInput + ". You are " + ageInput + " years old.")
}

/*
Notes:
1. `bufio.NewReader(os.Stdin)` is used to read user input.
2. `ReadString('\n')` reads input until a newline is encountered.
3. `strings.TrimSpace(input)` removes the trailing newline and extra spaces.
4. The `_` (blank identifier) is used to ignore errors (comma ok idiom).
5. Always sanitize input using `TrimSpace()` to avoid unwanted line breaks.
*/
