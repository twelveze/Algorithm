package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	var minPrice int
	var maxProfit int
	minPrice = math.MaxInt
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice >= maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}
