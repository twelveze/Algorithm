package main

import (
	"Algorithm/Algorithmtest/LeetCodeDaily/common"
	"fmt"
)

//https://leetcode.cn/problems/longest-ideal-subsequence/
//思路：暴力DP
const MAXN = 1e5 + 50

func longestIdealString(s string, k int) int {
	len := len(s)
	dp := make([][]int, MAXN) //dp[i][j]表示s的前i个字母中的以j结尾的理想字符串的最长长度
	for i, _ := range dp {
		dp[i] = make([]int, 30)
	}
	res := 0
	for i := 1; i <= len; i++ {
		c := int(s[i-1] - 'a')
		for j := 0; j < 26; j++ {
			dp[i][j] = dp[i-1][j]
		}
		for j := 0; j < 26; j++ {
			if common.Abs(c, j) <= k {
				dp[i][c] = common.Max(dp[i][c], dp[i-1][j]+1)
			}
		}
	}
	for i := 0; i < 26; i++ {
		res = common.Max(res, dp[len][i])
	}
	return res
}
func main() {
	str := "acfgbd"
	k := 2
	res := longestIdealString(str, k)
	fmt.Println(res)
}
