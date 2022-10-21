package main

import (
	"fmt"
	"math"
)

type StockSpanner struct {
	stack [][2]int
	idx   int
}

func Constructor901() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt32}}, -1}
}

func (s *StockSpanner) Next(price int) int {
	s.idx++
	for price >= s.stack[len(s.stack)-1][1] {
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, [2]int{s.idx, price})
	return s.idx - s.stack[len(s.stack)-2][0]
}

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
