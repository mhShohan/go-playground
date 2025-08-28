package main

import (
	"strings"
)

/*
151. Reverse Words in a String
=> pseudocode - steps
- split the string by space
- iterate from the end of the array
- if the string is not empty then
	- if it's the first word then add it to the result
	- else add a space and then the word to the result
*/

func reverseWords(s string) string {
	str := strings.Split(s, " ") // split the string by space
	isFirst := true              // to track if it's the first word
	var result strings.Builder   // using strings.Builder for efficient string concatenation

	for i := len(str) - 1; i >= 0; i-- {
		if len(strings.TrimSpace(str[i])) > 0 {
			if !isFirst {
				result.WriteString(" ") // add space before the word if it's not the first
			}
			result.WriteString(str[i])
			isFirst = false
		}
	}

	return result.String()
}

// func main() {
// 	fmt.Println("-", reverseWords("the sky is blue"), "-")
// 	fmt.Println("-", reverseWords("  hello world  "), "-")
// 	fmt.Println("-", reverseWords("  hello        world"), "-")
// }
