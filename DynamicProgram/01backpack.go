package main

import (
	"Model"
	"fmt"
)

func backpack(weight, value []int, limit int) int {
	res := 0
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, limit+1)
	}
	for i := 0; i < len(weight); i++ {
		for j := 0; j <= limit; j++ {
			if i == 0 {
				dp[i][j] = 0
			} else if j == 0 {
				dp[i][j] = 0
			} else {
				if weight[i] > j {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = Model.Max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
				}
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}
func main() {
	limit := 8
	weight := []int{2, 3, 4, 5}
	value := []int{3, 4, 5, 8}
	res := backpack(weight, value, limit)
	fmt.Println(res)
}
