package main

import "fmt"

// leetcode -70 经典爬楼梯问题
func climbStepFunc1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStepFunc1(n-1) + climbStepFunc1(n-2)
}
func main() {
	n := 3
	fmt.Println(climbStepFunc1(n))
}
