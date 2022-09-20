package main

import (
	"Model"
	"fmt"
)

func longestPalindromeSubseq(s string) int {
	//dp[i][j] 代表s[i - j]这一段字符内的最长回文子序列
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	//每个字符对应的最长回文子序列是1(他本身)
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	//反着遍历来确保正确的遍历状态
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] { //首尾相等，可以作为一个回文
				dp[i][j] = dp[i+1][j-1] + 2
			} else { //首尾不想等，得看里面较小一层的max
				dp[i][j] = Model.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

//最长回文子序列
func main() {
	str := "abbba"
	res := longestPalindromeSubseq(str)
	fmt.Println(res)
}
