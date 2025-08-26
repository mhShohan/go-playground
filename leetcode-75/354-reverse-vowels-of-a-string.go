package main

var vowels = map[rune]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

func isVowel(r rune) bool {
	return vowels[r]
}

func reverseVowels(s string) string {
	string_runes := []rune(s)             // convert in Unicode
	left, right := 0, len(string_runes)-1 // left and right pointers

	for left < right {
		if isVowel(string_runes[left]) && isVowel(string_runes[right]) {
			string_runes[left], string_runes[right] = string_runes[right], string_runes[left] // ✅ swap
			left++
			right--
		} else {
			if !isVowel(string_runes[left]) {
				left++
			}
			if !isVowel(string_runes[right]) {
				right--
			}
		}
	}

	return string(string_runes)
}

// func main() {
// 	fmt.Println(reverseVowels("leetcode")) //leotcede
// }
