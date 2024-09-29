package main

import (
	"fmt"
	"strings"
)

func recursiveIncrement(N, n int) int {
	if N == n {
		return N
	}

	fmt.Println(N)

	return recursiveIncrement(N+1, n)
}

func recursiveDecrement(n int) int {
	if n == 0 {
		return 0
	}

	fmt.Println(n)

	return recursiveDecrement(n - 1)
}

func incrementBacktrack(i, n int) {
	if i < 1 {
		return
	}

	incrementBacktrack(i-1, n)
	fmt.Println(i)
}

func decrementBacktrack(i, n int) {
	if i > n {
		return
	}

	decrementBacktrack(i+1, n)
	fmt.Println(i)
}

func recursiveSum(n int) int {
	if n == 0 {
		return 0
	}

	return n + recursiveSum(n-1)
}

func recursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * recursiveFactorial(n-1)
}

// Reverse Array using Two Pointer
func reverseArray() {
	numbers := []int{10, 20, 30, 40, 50}

	arrayLength := len(numbers) - 1

	for i := 0; i <= arrayLength; i++ {
		start := i
		end := arrayLength - i

		if start >= end {
			break
		}

		numbers[start], numbers[end] = numbers[end], numbers[start]
	}

	fmt.Println(numbers)
}

// Reverse Array using Two Pointer and Recursion
func swap(start, end int, numbers []int) {
	if start >= end {
		return
	}

	numbers[start], numbers[end] = numbers[end], numbers[start]
	swap(start+1, end-1, numbers)
}

func reverseArrayTwoPointer() {
	numbers := []int{1, 2, 3, 4, 5}
	swap(0, len(numbers)-1, numbers)

	fmt.Println(numbers)
}

// reverse string using Two Pointer and Recursion
func reverseString(s string) string {
	str := []byte(s)
	reverse(0, len(str)-1, str)
	return string(str)
}

func reverse(start, end int, str []byte) {
	if start >= end {
		return
	}

	str[start], str[end] = str[end], str[start]
	reverse(start+1, end-1, str)
}

// check if a string is palindrome
// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// 97 - 122

func isPalindrome(s string) bool {
	str := strings.ToLower(s)

	return checkPalindrome(0, len(str)-1, str)
}

func checkPalindrome(start, end int, str string) bool {
	if start >= end {
		return true
	}

	if str[start] != str[end] {
		return false
	}

	return checkPalindrome(start+1, end-1, str)
}

func isAlphanumeric(char byte) bool {
	if (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') {
		return true
	}

	return false
}

func main() {
	fmt.Println(recursiveIncrement(1, 5))          // 1 2 3 4 5 (Start, End)
	fmt.Println(recursiveDecrement(5))             // 5 4 3 2 1
	fmt.Println(recursiveSum(5))                   // 15
	fmt.Println(recursiveFactorial(5))             // 120
	reverseArrayTwoPointer()                       // 5 4 3 2 1
	reverseArray()                                 // 50 40 30 20 10
	fmt.Println(reverseString("Hello"))            // olleH
	isPalindrome("A man, a plan, a canal: Panama") // true

	fmt.Println("Backtrack")
	incrementBacktrack(5, 5) // 5 4 3 2 1 (Start, End)
	decrementBacktrack(1, 5) // 1 2 3 4 5
}
