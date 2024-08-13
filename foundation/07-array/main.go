package main

import "fmt"

func main() {
	fmt.Println("Array")

	var names [10]string
	names[0] = "Mehdi"
	names[1] = "Hasan"
	names[2] = "Shohan"

	fmt.Println(names)
	fmt.Println("Length of names array: ", len(names))

	// Array with values
	ages := [3]int{10, 20, 30}
	fmt.Println(ages)
}
