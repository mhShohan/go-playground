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

	// taking input from the user using bufio package
	name := bufio.NewReader(os.Stdin)
	age := bufio.NewReader(os.Stdin)

	// comma ok idiom || _ is a blank identifier || err ok idiom
	nameInput, _ := name.ReadString('\n')
	ageInput, _ := age.ReadString('\n')

	fmt.Println("Hello, ", strings.TrimSpace(nameInput))

	// converting string to integer
	ageValue, err := strconv.ParseInt(strings.TrimSpace(ageInput), 10, 64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("You are ", ageValue, " years old")
	}
}
