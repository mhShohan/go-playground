package main

import (
	"fmt"
	"math"
)

func patternOne(n int) {
	fmt.Println("------------ Pattern 1 -----------------")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			print("*")
		}
		println()
	}
}

func patternTwo(n int) {
	fmt.Println("------------ Pattern 2 -----------------")

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			print("*")
		}
		println()
	}
}

func patternThree(n int) {
	fmt.Println("------------ Pattern 3 -----------------")

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			print(j)
		}
		println()
	}
}

func patternFour(n int) {
	fmt.Println("------------ Pattern 4 -----------------")

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			print(i)
		}
		println()
	}
}

func patternFive(n int) {
	fmt.Println("------------ Pattern 5 -----------------")

	for i := n; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			print("*")
		}
		println()
	}
}

func patternSix(n int) {
	fmt.Println("------------ Pattern 6 -----------------")

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			print(j)
		}
		println()
	}
}

func patternSeven(n int) {
	fmt.Println("------------ Pattern 7 -----------------")

	for i := 0; i < n; i++ {
		// space
		for j := 0; j < (n - i - 1); j++ {
			print(" ")
		}

		// stars
		for j := 0; j < (2*i + 1); j++ {
			print("*")
		}

		// space
		for j := 0; j < (n - i - 1); j++ {
			print(" ")
		}

		println()
	}
}

func patternEight(n int) {
	fmt.Println("------------ Pattern 8 -----------------")

	for i := n - 1; i >= 0; i-- {
		// space
		for j := 0; j < (n - i - 1); j++ {
			print(" ")
		}

		// stars
		for j := 0; j < (2*i + 1); j++ {
			print("*")
		}

		// space
		for j := 0; j < (n - i - 1); j++ {
			print(" ")
		}

		println()
	}
}

func patternNine(n int) {
	fmt.Println("------------ Pattern 9 -----------------")

	for i := 0; i < n; i++ {
		// space
		for j := 0; j < (n - i - 1); j++ {
			print(" ")
		}

		// stars
		for j := 0; j < (2*i + 1); j++ {
			print("*")
		}

		// space
		for j := 0; j < (n - i - 1); j++ {
			print(" ")
		}

		println()
	}

	for i := n - 1; i >= 0; i-- {
		// space
		for j := 0; j < (n - i - 1); j++ {
			print(" ")
		}

		// stars
		for j := 0; j < (2*i + 1); j++ {
			print("*")
		}

		// space
		for j := 0; j < (n - i - 1); j++ {
			print(" ")
		}

		println()
	}
}

func patternTen(n int) {
	fmt.Println("------------ Pattern 10 -----------------")

	for i := 1; i <= 2*n-1; i++ {
		stars := i
		if i > n {
			stars = 2*n - i
		}

		for j := 1; j <= stars; j++ {
			print("*")
		}

		println()
	}
}

func patternEleven(n int) {
	fmt.Println("------------ Pattern 11 -----------------")

	start := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			start = 1
		} else {
			start = 0
		}

		for j := 0; j <= i; j++ {
			print(start)
			start = 1 - start
		}

		println()
	}
}

func patternTwelve(n int) {
	fmt.Println("------------ Pattern 12 -----------------")

	for i := 1; i <= n; i++ {
		// up count number
		for j := 1; j <= i; j++ {
			print(j)
		}
		// space
		for j := 1; j <= (2 * (n - i)); j++ {
			print(" ")
		}

		// down count number
		for j := i; j >= 1; j-- {
			print(j)
		}

		println()
	}
}

func patternThirteen(n int) {
	fmt.Println("------------ Pattern 13 -----------------")

	count := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			print(count, " ")
			count++
		}

		println()
	}
}

func patternFourteen(n int) {
	fmt.Println("------------ Pattern 14 -----------------")

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%c", 'A'+j)
		}
		fmt.Println()
	}
}

func patternFifteen(n int) {
	fmt.Println("------------ Pattern 15 -----------------")

	for i := n; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf("%c", 'A'+j)
		}
		fmt.Println()
	}
}

func patternSixteen(n int) {
	fmt.Println("------------ Pattern 16 -----------------")

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", 'A'+i)
		}
		fmt.Println()
	}
}

func patternSevenTeen(n int) {
	fmt.Println("------------ Pattern 17 -----------------")

	for i := 0; i < n; i++ {
		// starting space
		for j := 1; j <= n-i-1; j++ {
			print(" ")
		}

		char := 'A'
		breakPoint := (2*i + 1) / 2

		for j := 1; j <= 2*i+1; j++ {
			fmt.Printf("%c", char)
			if j <= breakPoint {
				char++
			} else {
				char--
			}
		}

		// ending space
		for j := 1; j <= n-i-1; j++ {
			print(" ")
		}

		println()
	}
}

func patternEightTeen(n int) {
	fmt.Println("------------ Pattern 18 -----------------")

	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			fmt.Printf("%c", 'A'+n-j)
		}
		println()
	}
}

func patternNineTeen(n int) {
	fmt.Println("------------ Pattern 19 -----------------")

	for i := 0; i < n; i++ {

		// starting stars
		for j := 0; j < n-i; j++ {
			print("*")
		}

		// spaces
		for j := 0; j < 2*i; j++ {
			print(" ")
		}

		// starting stars
		for j := 0; j < n-i; j++ {
			print("*")
		}

		println()
	}

	for i := n - 1; i >= 0; i-- {

		// starting stars
		for j := 0; j < n-i; j++ {
			print("*")
		}

		// spaces
		for j := 0; j < 2*i; j++ {
			print(" ")
		}

		// starting stars
		for j := 0; j < n-i; j++ {
			print("*")
		}

		println()
	}
}

func patternTwenty(n int) {
	fmt.Println("------------ Pattern 20 -----------------")

	initSpace := 2*n - 2

	for i := 1; i <= 2*n-1; i++ {
		stars := i

		if i > n {
			stars = 2*n - i
		}

		// stars
		for j := 1; j <= stars; j++ {
			print("*")
		}

		// spaces
		for j := 1; j <= initSpace; j++ {
			print(" ")
		}

		// stars
		for j := 1; j <= stars; j++ {
			print("*")
		}

		println()

		if i < n {
			initSpace -= 2
		} else {
			initSpace += 2
		}
	}
}

func patternTwentyOne(n int) {
	fmt.Println("------------ Pattern 21 -----------------")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i == 0 || i == n-1 {
				print("* ")
			} else {
				if j == 0 || j == n-1 {
					print("* ")
				} else {
					print("  ")
				}
			}
		}

		println()
	}
}

func patternTwentyTwo(n int) {
	fmt.Println("------------ Pattern 22 -----------------")

	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			top := i
			left := j
			right := (2*n - 2) - j
			bottom := (2*n - 2) - i

			result := math.Min(math.Min(float64(top), float64(bottom)), math.Min(float64(left), float64(right)))

			print(n-int(result), " ")
		}

		println()
	}
}

// main function - the entry point
func main() {
	patternOne(5)
	patternTwo(5)
	patternThree(5)
	patternFour(5)
	patternFive(5)
	patternSix(5)
	patternSeven(5)
	patternEight(5)
	patternNine(5)
	patternTen(5)
	patternEleven(5)
	patternTwelve(5)
	patternThirteen(5)
	patternFourteen(5)
	patternFifteen(5)
	patternSixteen(5)
	patternSevenTeen(5)
	patternEightTeen(5)
	patternNineTeen(5)
	patternTwenty(5)
	patternTwentyOne(5)
	patternTwentyTwo(5)
}
