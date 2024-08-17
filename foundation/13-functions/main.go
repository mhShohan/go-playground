package main

import "fmt"

func main() {
	fmt.Println("--- Functions ---")

	result := sum(10, 20)
	fmt.Println("Result:", result)

	result = sumOfNumbers(10, 20, 30, 40, 50)
	fmt.Println("Result:", result)

	sum, sub := sumAndSubtract(100, 20)
	fmt.Println("Sum:", sum, "Subtract:", sub)

	// IFII
	func() {
		fmt.Println("IFII")
	}()
}

func sum(a int, b int) int {
	return a + b
}

func sumOfNumbers(a ...int) int {
	total := 0

	for _, value := range a {
		total += value
	}

	return total
}

// return multiple values
func sumAndSubtract(a int, b int) (int, int) {
	return a + b, a - b
}
