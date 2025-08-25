package main

/*
- Time complexity: O(n)
- Space complexity: O(n)
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	// append 0 into flowerbed at first & last of the main flowerbed array
	new_flowerbed := append([]int{0}, append(flowerbed, 0)...)

	// skip first & last in the loop
	for i := 1; i < len(new_flowerbed)-1; i++ {

		// check three consecutive empty plots to plant flower at the index: i
		if new_flowerbed[i-1] == 0 && new_flowerbed[i] == 0 && new_flowerbed[i+1] == 0 {
			new_flowerbed[i] = 1 // mark as plant
			n--                  // reduce the value of n
		}
	}

	return n <= 0
}

/*
- reduce the space complexity to O(1)
- Time complexity: O(n)
- Space complexity: O(1)
*/
func canPlaceFlowers2(flowerbed []int, n int) bool {
	for i := range flowerbed {
		// check flowerbed plot is empty
		if flowerbed[i] == 0 {
			// check item is first index or previous item is zero
			isLeftEmpty := (i == 0) || (flowerbed[i-1] == 0)

			// check item is the last index of next item is zero
			isRightEmpty := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)

			// check left and right side plot is empty
			if isLeftEmpty && isRightEmpty {
				flowerbed[i] = 1 // plan flower to flowerbed
				n--

				if n == 0 {
					return true
				}
			}
		}

	}

	return n <= 0
}

// func main() {
// 	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
// 	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
// }
