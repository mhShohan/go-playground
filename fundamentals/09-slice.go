package main

import "fmt"

func main() {
	numArr := []int{}
	mySlice := numArr[:]

	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, i)
	}

	fmt.Println(mySlice)
}
