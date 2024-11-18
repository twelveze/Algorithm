package main

import (
	"fmt"
	"model"
)

func maxProfit3(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = model.Max(buy1, -prices[i])
		sell1 = model.Max(sell1, buy1+prices[i])
		buy2 = model.Max(buy2, sell1-prices[i])
		sell2 = model.Max(sell2, buy2+prices[i])
	}
	return sell2
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	res := maxProfit3(prices)
	fmt.Println(res)
}
