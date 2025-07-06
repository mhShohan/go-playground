package main

import "fmt"

// counter returns two functions: increment and decrement.
// Both closures share the same 'count' variable which is local to counter().
func counter() (func() int, func() int) {
	count := 0 // This variable is captured by the returned closures.

	// increment increases the count and returns the updated value.
	increment := func() int {
		count += 1
		return count
	}

	// decrement decreases the count and returns the updated value.
	decrement := func() int {
		count -= 1
		return count
	}

	// Return both closures. Each closure has access to the 'count' variable.
	return increment, decrement
}

// init runs before main in Go. Useful for setup, configuration, or print banners.
func init() {
	fmt.Println("========Go init=========")
}

func main() {
	// Create the first counter instance
	inc1, dec1 := counter()

	fmt.Println("First counter instance (inc1/dec1):")
	fmt.Println("inc1():", inc1()) // 1
	fmt.Println("inc1():", inc1()) // 2
	fmt.Println("inc1():", inc1()) // 3
	fmt.Println("dec1():", dec1()) // 2
	fmt.Println("dec1():", dec1()) // 1

	fmt.Println()

	// Create a second independent counter instance
	inc2, dec2 := counter()

	fmt.Println("Second counter instance (inc2/dec2):")
	fmt.Println("inc2():", inc2()) // 1
	fmt.Println("inc2():", inc2()) // 2
	fmt.Println("inc2():", inc2()) // 3
	fmt.Println("dec2():", dec2()) // 2
	fmt.Println("dec2():", dec2()) // 1
	fmt.Println("dec2():", dec2()) // 0

	fmt.Println()

	// Reuse first instance to show it maintains its state independently
	fmt.Println("Back to first counter (inc1):")
	fmt.Println("inc1():", inc1()) // 2
	fmt.Println("inc1():", inc1()) // 3
}
