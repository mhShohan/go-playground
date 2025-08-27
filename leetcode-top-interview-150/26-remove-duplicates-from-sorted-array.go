package main

func removeDuplicates(nums []int) int {

	i := 0 // first pointer
	j := 1 // second pointer

	for j < len(nums) {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}

		j++
	}

	return i + 1
}
