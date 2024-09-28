package main

import "fmt"

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

func main() {
	fmt.Println(recursiveIncrement(1, 5)) // 1 2 3 4 5 (Start, End)
	fmt.Println(recursiveDecrement(5))    // 5 4 3 2 1
	fmt.Println(recursiveSum(5))          // 15
	fmt.Println("Backtrack")
	incrementBacktrack(5, 5) // 5 4 3 2 1 (Start, End)
	decrementBacktrack(1, 5) // 1 2 3 4 5
}
