package main

import (
	"fmt"
	"math"
)

func KeyBoarsWithTwoKeys(n int) int { //使用动态规划
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt64
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				dp[i] = MIN(dp[i], dp[j]+(i/j))
			}
		}
	}
	return dp[n]
}
func MIN(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func main() {
	n := 30
	res := KeyBoarsWithTwoKeys(n)
	fmt.Println(res)
}
