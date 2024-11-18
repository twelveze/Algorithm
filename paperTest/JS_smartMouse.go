package main

import "fmt"

func maxTotal(triangle [][]int) int {
	lens := len(triangle)
	dp := make([][]int, lens)
	for i := 0; i < lens; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < lens-1; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == i {
				dp[i][j] = dp[i][j-1] + triangle[i][j]
			} else {
				dp[i][j] = max2(dp[i-1][j], dp[i][j-1]) + triangle[i][j]
			}
		}
	}
	for i := 0; i < lens-1; i++ {
		dp[lens-1][i] = dp[lens-2][i] + triangle[lens-1][i]
	}
	res := 0
	for i := 0; i < lens-1; i++ {
		if dp[lens-1][i] > res {
			res = dp[lens-1][i]
		}
	}
	return res
}
func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	triangle := make([][]int, 0)
	for i := 1; i <= n; i++ {
		temp := make([]int, 0)
		var num int
		for j := 0; j < i; j++ {
			fmt.Scan(&num)
			temp = append(temp, num)
		}
		triangle = append(triangle, temp)
	}
	res := maxTotal(triangle)
	fmt.Println(res)
}
