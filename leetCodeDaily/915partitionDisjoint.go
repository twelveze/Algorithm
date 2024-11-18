package main

import (
	"fmt"
	"model"
)

func partitionDisjoint(nums []int) int {
	n := len(nums)
	leftMax := nums[0]
	leftPos := 0
	Max := nums[0]
	for i := 1; i < n-1; i++ {
		Max = model.Max(nums[i], Max)
		if leftMax > nums[i] {
			leftMax = Max
			leftPos = i
		}
	}
	return leftPos + 1
}

func main() {
	nums := []int{5, 0, 3, 8, 6}
	res := partitionDisjoint(nums)
	fmt.Println(res)
}
