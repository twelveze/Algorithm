package main

import (
	"fmt"
	"sort"
)

//https://leetcode.cn/problems/russian-doll-envelopes/
//https://labuladong.gitee.io/algo/3/25/70/
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool { //先对二维数组排序
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		} else {
			return envelopes[i][1] > envelopes[j][1]
		}
	})
	var nums []int
	for i := 0; i < len(envelopes); i++ {
		nums = append(nums, envelopes[i][1])
	}
	res := lengthOfLISForEnvelopes(nums)
	return res
}

//使用普通的LIS算法会超时,本题目hard难度一大部分在这一步的处理上
func lengthOfLISForEnvelopes(nums []int) int {
	f := []int{}
	for _, v := range nums {
		if i := sort.SearchInts(f, v); i < len(f) {
			f[i] = v
		} else {
			f = append(f, v)
		}
	}
	return len(f)
}
func main() {
	envelopes := [][]int{}
	row1 := []int{5, 4}
	row2 := []int{6, 4}
	row3 := []int{6, 7}
	row4 := []int{2, 3}
	envelopes = append(envelopes, row1, row2, row3, row4)
	res := maxEnvelopes(envelopes)
	fmt.Println(res)
}
