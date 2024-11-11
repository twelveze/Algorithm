package main

import (
	"Algorithm/Model"
	"fmt"
)

// leetcode-3 无重复字符的最长子串
func lengthOfLongestNoRepeatSubstring(str string) int {
	repeatMap := make(map[byte]int, 0)
	if len(str) == 0 {
		return 0
	}
	ans := 1
	//初始化窗口map
	repeatMap[str[0]] = 0
	for left, right := 0, 1; right < len(str); right++ {
		if _, ok := repeatMap[str[right]]; ok {
			rightChar := str[right]
			// 去掉旧窗口元素
			for i := left; i < repeatMap[rightChar]; i++ {
				delete(repeatMap, str[i])
			}
			// 更新窗口左边位置
			left = repeatMap[rightChar] + 1
			// 增加新窗口右边的元素
			repeatMap[rightChar] = right
		} else {
			repeatMap[str[right]] = right
		}
		// 每滑动一次计算一次
		length := right - left + 1
		ans = Model.Max(ans, length)
	}
	return ans
}
func main() {
	str := "abba"
	res := lengthOfLongestNoRepeatSubstring(str)
	fmt.Println(res)
}
