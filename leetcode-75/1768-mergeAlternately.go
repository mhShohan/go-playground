package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1, word2 string) string {
	var builder strings.Builder

	maxCount := len(word1)
	if len(word2) > maxCount {
		maxCount = len(word2)
	}

	for i := 0; i < maxCount; i++ {
		if i < len(word1) {
			builder.WriteByte(word1[i])
		}
		if i < len(word2) {
			builder.WriteByte(word2[i])
		}
	}

	return builder.String()
}

func main() {
	fmt.Println(mergeAlternately("ace", "bdfgh"))
}
