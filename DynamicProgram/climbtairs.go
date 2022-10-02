package main

import "fmt"

func climbStairs(n int) int{
	dp := make([]int,n + 1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++{
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}
func main() {
	var n = 6
	res := climbStairs(n)
	fmt.Println(res)
}
