package main

import "fmt"

func main() {
	myArr := []int{1, 2, 3, 4, 5}

	for i, v := range myArr {
		fmt.Println(i, v)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
