package main

func majorityElement(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
		if counts[num] > len(nums)/2 {
			return counts[num]
		}
	}

	return -1
}

// func main() {
// 	fmt.Println(majorityElement([]int{3, 2, 3}))
// 	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
// }
