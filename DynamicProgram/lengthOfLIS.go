package main

import "fmt"

func lengthOfLIS(nums []int) int{
	len := len(nums)
	res := 0
	dp := make([]int, len) //dp表示以第i个数字结尾的最长子序列长度
	for i := 0; i < len; i++{
		dp[i] = 1 //初始化为1
		for j := 0; j < i; j++{
			if nums[i] > nums[j] && dp[i] < dp[j] + 1{
				dp[i] = dp[j] + 1
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(a, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}
func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}