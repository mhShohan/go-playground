package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter Your Name: ")

	// Taking input from the user using bufio package
	reader := bufio.NewReader(os.Stdin)
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput) // Remove trailing newline

	fmt.Println("Enter Your Age: ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)

	fmt.Println("Hello, ", nameInput)

	// Converting string to integer
	ageValue, err := strconv.Atoi(ageInput)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("You are", ageValue, "years old")
	}
}

/*
Notes:
1. **Reading User Input:**
   - `bufio.NewReader(os.Stdin)` is used to create a buffered reader to take input from the user.
   - `ReadString('\n')` reads input until the user presses enter.

2. **Handling Newline Characters:**
   - When using `ReadString('\n')`, the input includes a newline character (`\n`).
   - `strings.TrimSpace(input)` is used to remove any leading or trailing whitespace, including `\n`.

3. **Handling Numeric Input:**
   - User input is always received as a string.
   - `strconv.Atoi(input)` converts the string input to an integer.
   - If the input is not a valid number, an error is returned.

4. **Error Handling:**
   - The `strconv.Atoi(input)` function can fail if the input is not a number.
   - The `if err != nil` block ensures the program does not crash and provides an error message.

5. **Code Efficiency:**
   - Using a single `bufio.Reader` instead of two separate readers optimizes memory usage.
   - `TrimSpace()` helps in preventing unwanted formatting issues due to extra spaces or newlines.
*/
