package main

import (
	"slices"
)

// kidsWithCandies determines which children can have the greatest number of candies if they are given some extra candies.

// Parameters:
//   - candies: a slice of integers where each element represents the number of candies a child currently has.
//   - extraCandies: the number of extra candies you can give to each child.

// Returns:
//   - A slice of booleans where each element corresponds to a child.
//     If the value is true, that child can have the greatest number of
//     candies after receiving the extra candies; otherwise false.

// Example:
//   candies = [2, 3, 5, 1, 3], extraCandies = 3
//   max = 5
//   - Child 1: 2+3 = 5 (>= max) → true
//   - Child 2: 3+3 = 6 (>= max) → true
//   - Child 3: 5+3 = 8 (>= max) → true
//   - Child 4: 1+3 = 4 (< max)  → false
//   - Child 5: 3+3 = 6 (>= max) → true
//   Result → [true, true, true, false, true]

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// Find the maximum number of candies any child currently has
	max := slices.Max(candies)
	output := []bool{}

	// Check each child
	for _, value := range candies {
		// If this child’s candies + extraCandies is at least equal to max,
		// then they can have the greatest number of candies
		if value+extraCandies >= max {
			output = append(output, true)
		} else {
			output = append(output, false)
		}
	}

	return output
}

/*
func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))  // [true true true false true]
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))  // [true false false false false]
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))     // [true false true]
}

*/
