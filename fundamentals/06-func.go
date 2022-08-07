package main

import "fmt"

func main() {
	fmt.Println(conut(1, 5))
	fmt.Println(conut(6, 10))
}

func conut(start int, end int) int {
	sum := 0

	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}
