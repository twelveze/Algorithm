package main

import (
	"fmt"
	"math"
	"model"
)

func maxSubArray(arr []int) int {
	min := 0
	ans := math.MinInt64
	num := 0
	for i := 0; i < len(arr); i++ {
		num += arr[i]
		ans = model.Max(ans, num-min)
		if num < min {
			min = num
		}
	}
	return ans
}
func maxSubArrayByDp(arr []int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	max := math.MinInt64
	for i := 1; i < len(arr); i++ {
		if arr[i]+dp[i-1] > arr[i] {
			dp[i] = arr[i] + dp[i-1]
		} else {
			dp[i] = arr[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
func main() {
	arr := []int{-8, 3, -2, 3}
	res := maxSubArray(arr)
	resByDp := maxSubArrayByDp(arr)
	fmt.Println(res)
	fmt.Println(resByDp)
}
