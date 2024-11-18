package main

import "fmt"

func maxValue(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i >= 1 {
			dp[i] = max1(dp[i-1]+3, dp[i])
		}
		if i >= 2 {
			dp[i] = max1(dp[i-2]+7, dp[i])
		}
		if i >= 3 {
			dp[i] = max1(dp[i-3]+11, dp[i])
		}
		if i >= 4 {
			dp[i] = max1(dp[i-4]+15, dp[i])
		}
		if i >= 5 {
			dp[i] = max1(dp[i-5]+20, dp[i])
		}
	}
	return dp[n]
}
func max1(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	res := maxValue(n)
	fmt.Println(res)
}
