package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("--- If Else ---")

	num := 15

	if num%3 == 0 {
		fmt.Println("Fiz")
	} else if num%5 == 0 {
		fmt.Println("Buz")
	} else if num%3 == 0 && num%5 == 0 {
		fmt.Println("FizBuz")
	} else {
		fmt.Println("Nothing")
	}

	if num := 10; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	fmt.Println("--- Switch Case ---")
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
}
