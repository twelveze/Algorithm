package main

import "fmt"

func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = maxx(buy1, -prices[i])
		sell1 = maxx(sell1, buy1+prices[i])
		buy2 = maxx(buy2, sell1-prices[i])
		sell2 = maxx(sell2, buy2+prices[i])
	}
	return sell2
}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}
