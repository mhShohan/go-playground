package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	// const s = "আদাফসদদফাসদফাসদফাদ"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
}
