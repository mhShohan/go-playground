package main

import "fmt"

func main() {
	fmt.Println("Structs")

	type User struct {
		name    string
		age     int
		city    string
		country string
	}

	user := User{name: "MH Shohan", age: 25, city: "Dhaka", country: "Bangladesh"}
	fmt.Println(user)
	fmt.Printf("%+v\n", user)
	fmt.Println("Name:", user.name)
	fmt.Println("Age:", user.age)

	// Updating struct
	user.age = 26
	fmt.Println("Updated Age:", user.age)
}
