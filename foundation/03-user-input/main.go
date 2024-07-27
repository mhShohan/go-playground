package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter Your Name: ")

	name := bufio.NewReader(os.Stdin)

	age := bufio.NewReader(os.Stdin)

	// comma ok idiom || _ is a blank identifier || err ok idiom
	nameInput, _ := name.ReadString('\n')
	ageInput, _ := age.ReadString('\n')

	fmt.Println("Hello, " + nameInput + ". You are " + ageInput + " years old")
}
