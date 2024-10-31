package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	// prices := []int{7, 6, 5, 4, 3, 2}
	// prices := []int{2, 4, 1}
	fmt.Println("profit", bestTimeToBuyAndSell(prices))
}

/*
121. Best Time to Buy and Sell Stock
- Initialize two variables `minPrice` and `profit` to the first element of the array
- Iterate over the array starting from the second element
- Calculate the `cost` as the difference between the current element and the `minPrice`
- Update the `profit` as the maximum of the `profit` and the `cost`
- Update the `minPrice` as the minimum of the `minPrice` and the current element
- Return the `profit`
*/
func bestTimeToBuyAndSell(prices []int) int {
	minPrice, profit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		cost := prices[i] - minPrice
		profit = max(profit, cost)
		minPrice = min(minPrice, prices[i])
	}

	return profit
}
