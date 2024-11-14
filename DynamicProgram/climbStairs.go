package main

import "fmt"

// leetcode -70 经典爬楼梯问题
func climbStepFunc2(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

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
