package main

import (
	"fmt"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	m := len(digits)
	s := strconv.Itoa(n)
	k := len(s)
	dp := make([][2]int, k+1)
	dp[0][1] = 1
	for i := 1; i <= k; i++ {
		for _, d := range digits {
			if d[0] == s[i-1] {
				dp[i][1] = dp[i-1][1]
			} else if d[0] < s[i-1] {
				dp[i][0] += dp[i-1][1]
			} else {
				break
			}
		}
		if i > 1 {
			dp[i][0] += m + dp[i-1][0]*m
		}
	}
	return dp[k][0] + dp[k][1]
}
func main() {
	digits := []string{"1", "3", "5", "7"}
	n := 100
	res := atMostNGivenDigitSet(digits, n)
	fmt.Println(res)
}
