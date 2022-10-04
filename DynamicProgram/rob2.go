package main

import (
	"Algorithm/Model"
	"fmt"
)

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Model.Max(nums[0], nums[1])
	}
	//这里要分成两个[]int来处理
	n := len(nums)
	num1 := nums[:n-1]
	num2 := nums[1:]
	res1 := robber(num1)
	res2 := robber(num2)
	return Model.Max(res1, res2)
}

func robber(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Model.Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = Model.Max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

//https://leetcode.cn/problems/house-robber-ii/
func main() {
	nums := []int{0, 0}
	res := rob2(nums)
	fmt.Println(res)
}
