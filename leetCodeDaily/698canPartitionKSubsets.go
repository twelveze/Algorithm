package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	var useMap = make(map[int]bool, len(nums))
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	per := sum / k
	sort.Ints(nums)
	n := len(nums)
	//最大的一位>平均也不行
	if nums[n-1] > per {
		return false
	}
	return backTrack(nums, len(nums)-1, per, 0, k, useMap)
}
func backTrack(nums []int, start int, target int, current int, k int, useMap map[int]bool) bool {
	if k == 1 { //// 分组操作执行k-1次之后，最后剩余的元素，就是最后一组了，不需要再匹配
		return true
	}
	if current == target { //这一组已经分好了,重制current
		return backTrack(nums, len(nums)-1, target, 0, k-1, useMap)
	}
	for i := start; i >= 0; i-- {
		if useMap[i] || nums[i]+current > target { // 被使用的元素，不能再次使用；总和大于目标值，也不能使用
			continue
		}
		useMap[i] = true // 这个数被选入分组
		if backTrack(nums, i-1, target, current+nums[i], k, useMap) {
			return true
		}
		useMap[i] = false
		for i > 0 && nums[i-1] == nums[i] { //// 例如“12333333...”，假如最右侧的“3”这个值没有匹配上，那么它左侧的剩余五个“3”都不需要再匹配了
			i--
		}
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 1, 2, 2, 2, 2}
	k := 2
	//回溯思想
	res := canPartitionKSubsets(nums, k)
	fmt.Println(res)
}
