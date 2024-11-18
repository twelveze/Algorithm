package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var temp []int
	sort.Ints(nums)
	//visit代表的是排序后第i位的数是否被纳入答案
	visit := make(map[int]bool, len(nums))
	ans = dfsForWholeSort2(ans, nums, temp, visit)
	return ans
}

func dfsForWholeSort2(ans [][]int, nums, temp []int, visit map[int]bool) [][]int {
	if len(temp) == len(nums) {
		copyTemp := make([]int, len(nums))
		copy(copyTemp, temp)
		ans = append(ans, copyTemp)
		return ans
	}
	for index := range nums {
		//紧挨着的两个数字一样，不能进入答案，因为前面必定已经有相同的组合了
		//这一步的判断是最关键的地方
		if index >= 1 && nums[index] == nums[index-1] && !visit[index-1] {
			continue
		}
		if !visit[index] {
			temp = append(temp, nums[index])
			visit[index] = true
			ans = dfsForWholeSort2(ans, nums, temp, visit)
			temp = temp[:len(temp)-1]
			visit[index] = false
		}
	}

	return ans
}
func main() {
	nums := []int{1, 1, 2}
	ans := permuteUnique(nums)
	fmt.Println(ans)
}
