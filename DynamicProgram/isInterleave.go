package main

import "fmt"

// leetcode - 97 交错字符串
/*
可以用动态规划来求解。我们定义 f(i,j) 表示 s的前 i 个元素和 s的前 j 个元素是否能交错组成 s的前 i+j 个元素。
如果s1的第 i 个元素和s3的第 i+j 个元素相等，那么s1的前 i 个元素和s2的前 j 个元素是否能交错组成s3的前 i+j 个元素取决于
s1的前 i−1 个元素和s2的前 j 个元素是否能交错组成s3的前 i+j−1 个元素，即此时 f(i,j) 取决于 f(i−1,j)，在此情况下如果 f(i−1,j) 为真，
则 f(i,j) 也为真。同样的，如果s2的第 j 个元素和s3的第 i+j 个元素相等并且 f(i,j−1) 为真，则f(i,j) 也为真。
*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	s1len, s2len, s3len := len(s1), len(s2), len(s3)
	if s1len+s2len != s3len {
		return false
	}
	//初始化状态转移方程
	dp := make([][]bool, s1len+1)
	for i := 0; i <= s1len; i++ {
		dp[i] = make([]bool, s2len+1)
	}
	dp[0][0] = true
	for i := 0; i <= s1len; i++ {
		for j := 0; j <= s2len; j++ {
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[s1len][s2len]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
}
