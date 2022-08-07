package main

import "fmt"

func main() {
	slice := make([]int, 1)

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)

	//copy the slice to another variable
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)

	fmt.Println(copySlice)

	//declarative slice
	declSl := []int{0, 1, 2, 3}
	fmt.Println(declSl)

}
