package main

import "fmt"

// User struct representing a user with name, age, and admin status
type User struct {
	Name    string
	Age     int
	isAdmin bool
}

// Method to get user information
func (user User) GetInfo() string {
	res := fmt.Sprintf("Name: %v, Age: %v, isAdmin: %v", user.Name, user.Age, user.isAdmin)
	return res
}

// Method to set user name (uses pointer receiver to modify original struct)
func (user *User) SetName(name string) {
	user.Name = name
}

// Method to check if user is an admin
func (user User) IsAdmin() bool {
	return user.isAdmin
}

// Method to update age
func (user *User) SetAge(age int) {
	user.Age = age
}

func main() {
	fmt.Println("--------Methods---------")

	// Creating a new User instance
	user := User{"mhShohan", 25, true}

	// Getter method
	result := user.GetInfo()
	fmt.Println(result)

	// Setter method - updating name
	user.SetName("MH Shohan")
	fmt.Println(user.GetInfo())

	// Checking if the user is an admin
	if user.IsAdmin() {
		fmt.Println("User is an admin.")
	} else {
		fmt.Println("User is not an admin.")
	}

	// Updating user age
	user.SetAge(30)
	fmt.Println("Updated User Info:", user.GetInfo())
}

/*
Notes:
=========
- **Methods in Go**
  - Methods are functions with a receiver (struct or pointer to struct).
  - Value receivers (`func (t Type) Method()`) don’t modify the original struct.
  - Pointer receivers (`func (t *Type) Method()`) allow modification of struct fields.
  - Methods help encapsulate behavior related to structs, making code more readable and modular.

- **Pointers in Methods**
  - Use pointer receivers when modifying struct fields to avoid copying the entire struct.
  - Pointer methods ensure changes persist outside the function scope.

Key Takeaways:
- Use value receivers for read-only operations.
- Use pointer receivers when modifying struct fields.
- Methods improve struct usability and keep related behavior together.
*/
