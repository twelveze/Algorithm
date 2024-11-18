package main

import (
	"fmt"
	"model"
)

func jump1(nums []int) int {
	pos := len(nums) - 1
	var step int
	for pos > 0 {
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos {
				pos = i
				step++
				break
			}
		}
	}
	return step
}

//这个方法很精妙
func jump2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	if len(nums) == 2 {
		return 1
	}
	length := len(nums)
	//只需要走到最后一个位置的前一位就行了，因为最少也要从前一位跳过去
	maxPos := 0    //maxPOs是下一次跳跃可以到达的边界
	end := nums[0] //end是本次跳跃可以到达的边界
	step := 1      //至少要一步
	for i := 1; i < length-1; i++ {
		maxPos = model.Max(maxPos, i+nums[i])
		if i == end {
			end = maxPos
			step++
		}
	}
	return step
}

func main() {
	nums := []int{1, 2, 3}
	res1 := jump1(nums)
	res2 := jump2(nums)
	fmt.Println(res1)
	fmt.Println(res2)
}
