package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, length) //dp[i][j]代表[i][j]处的最小值
	for i, arr := range triangle {
		dp[i] = make([]int, len(arr))
	}
	for i := 0; i < len(triangle[length-1]); i++ {
		dp[length-1][i] = triangle[length-1][i] //初始化最后一层dp
	}
	for i := length - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j]+triangle[i][j], dp[i+1][j+1]+triangle[i][j])
		}
	}
	return dp[0][0]
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res := minimumTotal(triangle)
	fmt.Println(res)
}
