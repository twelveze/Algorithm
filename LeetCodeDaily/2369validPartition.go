package main

import "fmt"

const maxn = 1e5 + 50

//https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/
func validPartition(nums []int) bool {
	dp := make([]bool, maxn) //表示nums的前n位是有效的
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		dp[i] = false
		//枚举三种情况
		if i >= 2 && nums[i-1] == nums[i-2] {
			dp[i] = dp[i] || dp[i-2]
		}
		if i >= 3 && nums[i-1] == nums[i-3] && nums[i-2] == nums[i-3] {
			dp[i] = dp[i] || dp[i-3]
		}
		if i >= 3 && nums[i-2] == nums[i-3]+1 && nums[i-1] == nums[i-2]+1 {
			dp[i] = dp[i] || dp[i-3]
		}
	}
	return dp[len(nums)]
}

func main() {
	nums := []int{4, 4, 4, 5, 6}
	res := validPartition(nums)
	fmt.Println(res)
}
