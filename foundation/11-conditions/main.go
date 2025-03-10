package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("--- If Else ---")

	// Example 1: Basic If-Else
	num := 15

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("Nothing")
	}

	// Example 2: Short Variable Declaration in If Condition
	if num := 10; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Example 3: Checking Ranges
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	fmt.Println("--- Switch Case ---")

	// Example 1: Rolling a Dice
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(6) + 1

	switch randNum {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six and roll again")
	default:
		fmt.Println("????")
	}

	// Example 2: Multiple Cases in One Condition
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Example 3: Switch without an Expression (if-else alternative)
	n := -5
	switch {
	case n > 0:
		fmt.Println("Positive Number")
	case n < 0:
		fmt.Println("Negative Number")
	default:
		fmt.Println("Zero")
	}
}

/*
Notes:
=========
- **If-Else Statements**
  - Used for conditional branching.
  - Short variable declaration is possible (`if var := value; condition {}`).
  - Multiple conditions can be checked using `else if`.
  - Logical operators (`&&`, `||`, `!`) help in combining conditions.

- **Switch Statements**
  - Useful for checking multiple conditions more cleanly than `if-else` chains.
  - Can be used with values (`switch variable {}`) or as an alternative to if-else (`switch {}` without an expression).
  - Multiple cases can share the same block (`case A, B:`).
  - `default` executes when no case matches.

- **Random Numbers & Time Package**
  - `rand.Seed(time.Now().UnixNano())` initializes the random number generator.
  - `rand.Intn(N)` generates a random integer from `0` to `N-1`.
  - `time.Now().Weekday()` gets the current day of the week.

Key Takeaways:
- Prefer `switch` over multiple `if-else` for readability.
- Always seed `rand` when generating random numbers to avoid repeated sequences.
- Use short variable declarations in `if` conditions to limit scope.
*/
