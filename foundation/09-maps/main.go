package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := make(map[string]string)

	user["name"] = "MH Shohan"
	user["age"] = "25"
	user["city"] = "Dhaka"
	user["country"] = "Bangladesh"

	fmt.Println(user)
	fmt.Println("Name:", user["name"])

	delete(user, "age")
	fmt.Println(user)

	// Iterating over map
	for key, value := range user {
		fmt.Printf("%s: %s\n", key, value)
	}
}
