package main

import (
	"fmt"
)

// StockSpanner 本质上是两个stack,这里使用切片模仿
type StockSpanner struct {
	prices []int
	weight []int
}

//初始化
func Constructor901() StockSpanner {
	return StockSpanner{[]int{}, []int{}}
}

func (s *StockSpanner) Next(price int) int {
	weight := 1
	//新加入一个price,如果大于等于栈顶的元素,则需要更新栈
	for len(s.prices) > 0 && s.prices[len(s.prices)-1] <= price {
		s.prices = s.prices[:len(s.prices)-1] //pop操作
		weight += s.weight[len(s.weight)-1]
		s.weight = s.weight[:len(s.weight)-1] //pop操作
	}
	s.prices = append(s.prices, price)
	s.weight = append(s.weight, weight)
	return weight
}

//https://leetcode.cn/problems/online-stock-span/solutions/18776/gu-piao-jie-ge-kua-du-by-leetcode/
//单调栈,更通俗易懂
func main() {
	stockSpanner := Constructor901()
	fmt.Println(stockSpanner.Next(100))
	fmt.Println(stockSpanner.Next(80))
	fmt.Println(stockSpanner.Next(60))
	fmt.Println(stockSpanner.Next(70))
	fmt.Println(stockSpanner.Next(60))
	fmt.Println(stockSpanner.Next(75))
	fmt.Println(stockSpanner.Next(85))
}
