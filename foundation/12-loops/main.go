package main

import "fmt"

func main() {
	fmt.Println("--- For Loop ---")
	fmt.Println("---------------------------------------------------------")

	users := []string{"John", "Doe", "Smith"}

	// Example 1: Traditional For Loop
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	fmt.Println("---------------------------------------------------------")

	// Example 2: For Loop with range (Index Only)
	for i := range users {
		fmt.Println("Index:", i, "Value:", users[i])
	}

	fmt.Println("---------------------------------------------------------")

	// Example 3: For Loop with range (Index and Value)
	for index, value := range users {
		fmt.Println("Index:", index, "Value:", value)
	}

	fmt.Println("-------------------- While Loop (For Alternative) -------------------------------------")

	// Example 4: While Loop Alternative in Go
	i := 0
	for i < len(users) {
		fmt.Println(users[i])
		i++
	}

	fmt.Println("-------------------- Infinite Loop with Break -------------------------------------")

	// Example 5: Infinite Loop with Break Condition
	j := 0
	for {
		if j >= len(users) {
			break
		}
		fmt.Println("User:", users[j])
		j++
	}

	fmt.Println("-------------------- Continue Statement -------------------------------------")

	// Example 6: Using Continue to Skip an Iteration
	for k := 0; k < len(users); k++ {
		if users[k] == "Doe" {
			continue
		}
		fmt.Println("User:", users[k])
	}
}

/*
Notes:
=========
- **For Loops in Go**
  - The only looping construct in Go is `for` (no `while` or `do-while`).
  - The `for` loop can act as a traditional loop (`for init; condition; post {}`) or as a while-loop alternative (`for condition {}`).
  - The `range` keyword simplifies iteration over slices, arrays, maps, and strings.
  - The index and value can both be captured in `range`, or just the index if the value is not needed.

- **Loop Control Statements**
  - `break` exits a loop immediately.
  - `continue` skips the rest of the current iteration and moves to the next one.
  - Infinite loops can be created using `for {}` and must be exited with `break`.

Key Takeaways:
- Prefer `range` for iterating over collections.
- Use `break` and `continue` to control loop flow efficiently.
- Go does not have `while`, but `for` can replicate it.
*/
