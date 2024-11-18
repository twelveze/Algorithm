package main

import (
	"fmt"
	"model"
)

/* leetcode -11 盛水最多的容器
*	#贪心 #双指针
 */

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		ans = model.Max(ans, (right-left)*model.MIN(height[left], height[right]))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return
}
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(height)
}
