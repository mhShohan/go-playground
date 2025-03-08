package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	// Creating a map using make
	user := make(map[string]string)

	// Adding key-value pairs to the map
	user["name"] = "MH Shohan"
	user["age"] = "25"
	user["city"] = "Dhaka"
	user["country"] = "Bangladesh"

	fmt.Println("User Map:", user)
	fmt.Println("Name:", user["name"])

	// Deleting a key from the map
	delete(user, "age")
	fmt.Println("After deletion:", user)

	// Iterating over the map using range
	fmt.Println("\nIterating over map:")
	for key, value := range user {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Checking if a key exists in the map
	if val, exists := user["age"]; exists {
		fmt.Println("Age exists with value:", val)
	} else {
		fmt.Println("Age key does not exist")
	}

	// Creating and initializing a map in one step
	person := map[string]int{
		"John": 30,
		"Doe":  25,
		"Jane": 28,
	}

	fmt.Println("\nPerson Age Map:", person)

	// Updating a map value
	person["Doe"] = 26
	fmt.Println("Updated Person Map:", person)

	// Finding length of the map
	fmt.Println("Number of entries in person map:", len(person))
}

/*
Notes:
1. **What is a Map?**
   - A map is a built-in data type that associates keys with values.
   - Syntax: `map[KeyType]ValueType{}`
   - Example: `user := make(map[string]string)`

2. **Adding and Accessing Elements**
   - Add elements: `mapVar[key] = value`
   - Retrieve values: `mapVar[key]`

3. **Deleting Elements**
   - Use `delete(mapVar, key)` to remove a key-value pair.

4. **Checking if a Key Exists**
   - Use `val, exists := mapVar[key]` to check if a key is present.
   - If `exists` is `false`, the key does not exist.

5. **Iterating Over a Map**
   - Use `for key, value := range mapVar {}` to loop through all key-value pairs.

6. **Initializing a Map with Values**
   - Example: `person := map[string]int{"John": 30, "Doe": 25}`

7. **Updating Values**
   - Simply assign a new value: `mapVar[key] = newValue`

8. **Getting the Length of a Map**
   - Use `len(mapVar)` to get the number of key-value pairs.

9. **Key Characteristics of Maps**
   - Unordered: The order of iteration is not guaranteed.
   - Dynamic: Maps grow automatically when new key-value pairs are added.
   - Fast: Lookup, insert, and delete operations have an average O(1) time complexity.
*/
