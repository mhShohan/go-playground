package main

import (
	"fmt"
)

// Note: In a real program, this would be in a separate package
type Person struct {
	Name    string // Capitalized = exported (public)
	Age     int    // Capitalized = exported (public)
	address string // Lowercase = unexported (private)
}

func main() {
	fmt.Println("=== STRUCTS IN GO ===")
	fmt.Println()

	// Basic struct declaration and initialization
	fmt.Println("1. Basic Struct Usage")
	fmt.Println("---------------------")

	// Define a struct type with fields
	type User struct {
		name    string
		age     int
		city    string
		country string
	}

	// Create a struct instance with field:value syntax
	user := User{name: "MH Shohan", age: 25, city: "Dhaka", country: "Bangladesh"}

	// Different ways to print struct data
	fmt.Println(user)         // Basic print
	fmt.Printf("%+v\n", user) // Print with field names
	fmt.Printf("%#v\n", user) // Go syntax representation

	// Accessing individual fields with dot notation
	fmt.Println("Name:", user.name)
	fmt.Println("Age:", user.age)

	// Updating struct fields
	user.age = 26
	fmt.Println("Updated Age:", user.age)
	fmt.Println()

	// Alternative initialization methods
	fmt.Println("2. Different Initialization Methods")
	fmt.Println("----------------------------------")

	// Method 1: Positional initialization (not recommended as it's brittle)
	user1 := User{"John Doe", 30, "New York", "USA"}
	fmt.Println("User1 (positional):", user1)

	// Method 2: Named fields (preferred)
	user2 := User{
		name:    "Alice Smith",
		age:     28,
		city:    "London",
		country: "UK",
	}
	fmt.Println("User2 (named fields):", user2)

	// Method 3: Create empty struct and assign fields later
	var user3 User
	user3.name = "Bob Johnson"
	user3.age = 35
	user3.city = "Sydney"
	user3.country = "Australia"
	fmt.Println("User3 (assigned after creation):", user3)
	fmt.Println()

	// Nested structs
	fmt.Println("3. Nested Structs")
	fmt.Println("-----------------")

	// Define an Address struct
	type Address struct {
		street  string
		city    string
		country string
		zipCode string
	}

	// Define an enhanced User struct with nested Address
	type EnhancedUser struct {
		name    string
		age     int
		address Address
		active  bool
	}

	// Create instance with nested struct
	enhancedUser := EnhancedUser{
		name: "Maria Garcia",
		age:  32,
		address: Address{
			street:  "123 Main St",
			city:    "Barcelona",
			country: "Spain",
			zipCode: "08001",
		},
		active: true,
	}

	fmt.Printf("Enhanced User: %+v\n", enhancedUser)
	fmt.Println("Address:", enhancedUser.address.city+", "+enhancedUser.address.country)
	fmt.Println()

	// Anonymous structs
	fmt.Println("4. Anonymous Structs")
	fmt.Println("-------------------")

	// Create a one-time use struct without defining a type
	temporaryUser := struct {
		name  string
		email string
		role  string
	}{
		name:  "Guest User",
		email: "guest@example.com",
		role:  "visitor",
	}

	fmt.Printf("Temporary User: %+v\n", temporaryUser)
	fmt.Println()

	// Struct with exported fields
	fmt.Println("5. Exported Fields (Public/Private)")
	fmt.Println("----------------------------------")

	// Note: In a real program, this would be in a separate package
	type Person struct {
		Name    string // Capitalized = exported (public)
		Age     int    // Capitalized = exported (public)
		address string // Lowercase = unexported (private)
	}

	person := Person{
		Name:    "Public Name",
		Age:     40,
		address: "Private Address",
	}

	fmt.Printf("Person: %+v\n", person)
	fmt.Println("Public field:", person.Name)
	fmt.Println("Private field (only accessible within package):", person.address)
	fmt.Println()

	// Struct methods
	fmt.Println("6. Methods on Structs")
	fmt.Println("--------------------")

	// Example using the Person struct
	// In Go, methods are defined outside the struct
	// displayPerson(person)

	// Create a new person with the constructor-like function
	newPerson := newPerson("Created Person", 50)
	fmt.Printf("New person: %+v\n", newPerson)
}

// A function that acts on a Person (not a method)
func displayPerson(p Person) {
	fmt.Printf("Person Display: %s is %d years old\n", p.Name, p.Age)
}

// A "constructor-like" function that creates and returns a Person
func newPerson(name string, age int) Person {
	return Person{
		Name:    name,
		Age:     age,
		address: "Default Address", // Private field set by constructor
	}
}

/*
==============================
NOTES: STRUCTS IN GO
==============================

Definition:
A struct in Go is a user-defined type that represents a collection of fields.
It's similar to classes in object-oriented programming languages but without
inheritance.

Key Concepts:

1. Declaration:
   - Structs are defined with the 'type' keyword followed by the struct name and fields
   - Fields have a name and a type
   - Example: type User struct { name string, age int }

2. Initialization:
   - Named fields: user := User{name: "John", age: 30}
   - Positional: user := User{"John", 30} (not recommended)
   - Empty struct: var user User (fields get zero values)
   - You can also create an instance and assign values later

3. Access and Modification:
   - Access fields with dot notation: user.name, user.age
   - Modify fields with assignment: user.age = 31

4. Zero Values:
   - Numeric fields: 0
   - String fields: ""
   - Boolean fields: false
   - Pointer fields: nil

5. Field Visibility:
   - Capitalized fields (Name) are exported (public)
   - Lowercase fields (name) are unexported (private)
   - Exported fields are accessible from other packages

6. Nested Structs:
   - Structs can contain other structs as fields
   - Access nested fields with chained dot notation: user.address.city

7. Anonymous Structs:
   - Can be created without defining a type
   - Useful for one-time use structures

8. Methods:
   - Go doesn't have classes, but you can define methods on structs
   - Methods are functions with a receiver argument
   - Example: func (u User) GetFullName() string { return u.firstName + " " + u.lastName }

Common Patterns:

1. Constructor Functions:
   - Go doesn't have constructors, but functions that return structs serve this purpose
   - Example: func NewUser(name string, age int) User { return User{name: name, age: age} }

2. Composition over Inheritance:
   - Go uses struct embedding for composition
   - Example: type Admin struct { User, permissions []string }

3. Struct Tags:
   - Metadata attached to struct fields
   - Used by libraries for reflection
   - Example: type User struct { Name string `json:"name"` }

Best Practices:

1. Use named fields for clarity, especially in larger structs
2. Keep structs focused on a single responsibility
3. Use proper capitalization to control visibility
4. Consider using constructor functions for complex initialization
5. Prefer composition over inheritance

Memory Considerations:
- Structs are value types in Go (passed by value, not by reference)
- Large structs should be passed as pointers to avoid copying
- Example: func modifyUser(u *User) { u.age = 31 }

*/
