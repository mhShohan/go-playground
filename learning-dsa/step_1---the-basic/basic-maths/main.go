package main

import (
	"fmt"
	"math"
	"sort"
)

func countDigit(n int) {
	N := n
	count := 0

	for N > 0 {
		// lastDigit:= N % 10
		N = int(N / 10)

		count++
	}

	fmt.Println(count)
}

func countDigitOptimized(n float64) {
	logValue := math.Log10(n)
	result := int(logValue + 1)
	fmt.Println(result)
}

func reverseNumber(n int) {
	N := n
	reverseNum := 0
	sign := 1

	if N < 0 {
		sign = -1
	}

	N *= sign

	for N != 0 {
		lastDigit := N % 10
		N = int(N / 10)

		reverseNum = (reverseNum * 10) + lastDigit
	}

	fmt.Println(reverseNum * sign)
}

func checkPalindrome(n int) {
	N := n
	reverseNumber := 0

	for N > 0 {
		lastDigit := N % 10
		N = int(N / 10)

		reverseNumber = (reverseNumber * 10) + lastDigit
	}

	fmt.Println(reverseNumber == n)
}

func armstrongNumber(n int) {
	N := n
	result := 0
	for N > 0 {
		lastDigit := N % 10
		N = int(N / 10)

		result += lastDigit * lastDigit * lastDigit
	}

	fmt.Println(n, "=>", result == n)
}

func printAllDivisors(n int) {
	divisors := []int{}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	fmt.Println(divisors)
	// time complexity: O(n)
}

func printAllDivisorsOptimized(n int) {
	divisors := []int{}

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)

			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}

	sort.Ints(divisors)

	fmt.Println(divisors)

	// time complexity: O(sqrt(n))
}

func main() {
	countDigit(23453345)
	countDigitOptimized(23453345)
	reverseNumber(-12345)
	reverseNumber(12345)

	checkPalindrome(12321)
	checkPalindrome(124241)
	armstrongNumber(371)
	armstrongNumber(5434)
	printAllDivisors(24)
	printAllDivisorsOptimized(36)
}
