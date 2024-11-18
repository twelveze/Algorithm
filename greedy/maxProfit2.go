package main

import (
	"fmt"
)

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var minPrice int
	var maxProfit int
	minPrice = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > minPrice && prices[i] >= prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		} else if prices[i] < prices[i-1] {
			minPrice = prices[i]
		} else if prices[i] > minPrice {
			maxProfit += prices[i] - minPrice
		}
	}
	return maxProfit
}
func main() {
	prices := []int{5, 2, 3, 2, 6, 6, 2, 9, 1, 0, 7, 4, 5, 0}
	res := maxProfit2(prices)
	fmt.Println(res)
}
