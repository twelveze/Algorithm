package main

import (
	"Algorithm/Model"
	"fmt"
)

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	for left < right {
		leftMax = Model.Max(height[left], leftMax)
		rightMax = Model.Max(height[right], rightMax)
		if height[left] < height[right] {
			ans = ans + (leftMax - height[left])
			left++
		} else {
			ans = ans + (rightMax - height[right])
			right--
		}
	}
	return ans
}

//接雨水
//https://leetcode.cn/problems/trapping-rain-water/
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}
