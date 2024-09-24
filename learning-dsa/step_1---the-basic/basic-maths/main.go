package main

import (
	"fmt"
	"math"
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

func main() {
	countDigit(23453345)
	countDigitOptimized(23453345)
	reverseNumber(-12345)
	reverseNumber(12345)

}
