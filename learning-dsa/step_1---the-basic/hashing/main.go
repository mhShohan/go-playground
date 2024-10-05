package main

import "fmt"

/*
The function `countNumberUsingHashing` takes an array of integers and returns a map with each
integer as a key and its frequency as the value.
@param {[]int} arr - The `arr` parameter in the `countNumberUsingHashing` function is a slice of
integers. This function takes this slice as input and returns a map where the keys are the unique
numbers from the input slice and the values are the counts of each number in the input slice.
*/
func countNumberUsingHashing(arr []int) map[int]int {
	count := make(map[int]int)

	for _, value := range arr {
		count[value]++
	}

	return count
}

/*
The function `characterCountUsingHashing` takes a string as input and returns a map with the count
of each character in the string using hashing.
@param {string} str - The `characterCountUsingHashing` function takes a string `str` as input and
returns a map where the keys are individual characters in the input string and the values are the
count of each character in the string.
*/
func characterCountUsingHashing(str string) map[string]int {
	count := make(map[string]int)

	for _, value := range str {
		count[string(value)]++
	}

	return count
}

/*
The function characterHashing calculates the frequency of a specific character in a given string.
@param {string} str - The `str` parameter in the `characterHashing` function is a string that
represents the input string for which we want to count the occurrences of a specific character.
@param {byte} character - The `characterHashing` function takes a string `str` and a byte
`character` as input parameters. The function calculates the frequency of each character in the
input string `str` and returns the count of the specified `character` in the string.
@returns The function `characterHashing` returns the count of occurrences of a specific character in
the input string.
*/
func characterHashing(str string, character byte) int {
	count := [256]int{}

	for i := 0; i < len(str); i++ {
		count[str[i]]++
	}

	return count[character]
}

/*
The function `findTheFrequencyOfNumber` takes an array of integers and returns a map with the
frequency of each number in the array.
@param {[]int} arr - The `findTheFrequencyOfNumber` function takes in a slice of integers `arr` as
input. It calculates the frequency of each number in the input slice and returns a map where the
keys represent the unique numbers in the input slice and the values represent the frequency of each
number in the input slice
@returns The function `findTheFrequencyOfNumber` returns a map where the keys are the unique numbers
*/
func findTheFrequencyOfNumber(arr []int) map[int]int {
	frequency := make(map[int]int)

	for _, value := range arr {
		frequency[value]++
	}

	return frequency
}

/*
The function `findHighestAndLowestFrequencyElement` takes an array of integers and returns the
highest and lowest frequency elements along with their frequencies.
@param {[]int} arr - The `findHighestAndLowestFrequencyElement` function you provided takes an
integer array `arr` as input and returns the highest and lowest frequency elements in the array.
@returns The function `findHighestAndLowestFrequencyElement` returns the highest and lowest
*/
func findHighestAndLowestFrequencyElement(arr []int) (int, int) {
	frequency := make(map[int]int)

	for _, value := range arr {
		frequency[value]++
	}

	highestFrequency := 0
	lowestFrequency := len(arr)

	var highestFrequencyElement, lowestFrequencyElement int

	for key, value := range frequency {
		if value > highestFrequency {
			highestFrequency = value
			highestFrequencyElement = key
		}

		if value < lowestFrequency {
			lowestFrequency = value
			lowestFrequencyElement = key
		}
	}

	fmt.Println("Highest Frequency element of array", highestFrequencyElement)
	fmt.Println("Lowest Frequency element of array", lowestFrequencyElement)

	return highestFrequency, lowestFrequency
}

func main() {
	fmt.Println("Hashing")
	arr := []int{1, 3, 4, 1, 2, 4, 1, 2, 3, 4, 11, 11, 12, 11, 11}

	fmt.Println(countNumberUsingHashing(arr))

	str := "hello! mehdi hasan shohan, how are you?"
	fmt.Println(characterCountUsingHashing(str))

	fmt.Println(characterHashing(str, 'h'))

	fmt.Println("Find the frequency of numbers", findTheFrequencyOfNumber(arr))
	findHighestAndLowestFrequencyElement(arr)
}
