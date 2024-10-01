package main

import "fmt"

/*
given array of number [1,3,4,1,2,4]
find the count of each number in the array
output: 1:2, 3:1, 4:2, 2:1
*/
func countNumberUsingHashing(arr []int) map[int]int {
	count := make(map[int]int)

	for _, value := range arr {
		count[value]++
	}

	return count
}

func characterCountUsingHashing(str string) map[string]int {
	count := make(map[string]int)

	for _, value := range str {
		count[string(value)]++
	}

	return count
}

func characterHashing(str string, character byte) int {
	count := [256]int{}

	for i := 0; i < len(str); i++ {
		count[str[i]]++
	}

	return count[character]
}

func main() {
	fmt.Println("Hashing")
	arr := []int{1, 3, 4, 1, 2, 4, 1, 2, 3, 4, 10, 11, 12, 11}
	fmt.Println(countNumberUsingHashing(arr))

	str := "hello! mehdi hasan shohan, how are you?"
	fmt.Println(characterCountUsingHashing(str))

	fmt.Println(characterHashing(str, 'h'))

}
