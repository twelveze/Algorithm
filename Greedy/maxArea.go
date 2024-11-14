package main

import (
	"Model"
	"fmt"
)

/* leetcode -11 盛水最多的容器
*	#贪心 #双指针
 */

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		ans = Model.Max(ans, (right-left)*Model.MIN(height[left], height[right]))
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
