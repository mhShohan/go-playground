package main

import "fmt"

func main() {
	fmt.Println("--- For Loop ---")
	fmt.Println("---------------------------------------------------------")

	users := []string{"John", "Doe", "Smith"}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	fmt.Println("---------------------------------------------------------")

	for i := range users {
		fmt.Println("Index:", i, "Value:", users[i])
	}

	fmt.Println("---------------------------------------------------------")

	for index, value := range users {
		fmt.Println("Index:", index, "Value:", value)
	}

	fmt.Println("--------------------while loop-------------------------------------")

	// while Loop
	i := 0
	for i < len(users) {
		fmt.Println(users[i])
		i++
	}
}
