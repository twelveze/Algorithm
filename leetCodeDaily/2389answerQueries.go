package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums, queries []int) []int {
	sort.Ints(nums)
	ans := make([]int, len(queries))
	tempNums := make([]int, len(nums)+1) //前缀和
	for i := 0; i < len(nums); i++ {
		tempNums[i+1] = tempNums[i] + nums[i]
	}
	for i := 0; i < len(ans); i++ {
		for j := 1; j <= len(nums); j++ {
			if queries[i] >= tempNums[j] {
				ans[i] = j
			}
		}
	}
	return ans
}

func main() {
	nums := []int{4, 5, 2, 1}
	queries := []int{3, 10, 21}
	res := answerQueries(nums, queries)
	fmt.Println(res)
}
