package main

import "fmt"

//接雨水
//https://leetcode.cn/problems/trapping-rain-water/
func PutRain(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans = 0
	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		//这种情况必定有leftMax < rightMax
		if height[left] < height[right] {
			ans = ans + (leftMax - height[left])
			left++
		} else { //这种情况必定有leftMax > rightMax
			ans = ans + (rightMax - height[right])
			right--
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := PutRain(height)
	fmt.Println(res)
}
