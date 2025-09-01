package main

import (
	"math"
)

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt32

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

// func main() {
// 	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
// 	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
// 	fmt.Println(maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
// }
