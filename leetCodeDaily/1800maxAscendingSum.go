package main

import "fmt"

func maxAscendingSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := 0
	for i := 0; i < len(nums)-1; i++ {
		tempSum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[j-1] {
				tempSum += nums[j]
			} else {
				break
			}
		}
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	return maxSum
}
func main() {
	nums := []int{11, 10}
	res := maxAscendingSum(nums)
	fmt.Println(res)
}
