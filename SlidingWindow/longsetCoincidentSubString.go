package main

import "fmt"

// huawei OD - 寻找符合要求的最长子串
/*
	给定一个字符串s，找出这样一个子串：
	1. 该子串中任意一个字符最多出现2次
	2. 该字串不包含指定某字符

	如果没有满足的子串则返回0，有则返回最长子串的长度
*/
func longestCoincidentSubString(s string, c byte) int {
	// 记录对应字符出现的次数
	charCount := make([]int, 127)
	ans := 0
	for left, right := 0, 0; right < len(s); right++ {
		char := s[right]
		// 遇到特殊字符直接从后一位进行计算
		if char == c {
			left = right + 1
			charCount = make([]int, 127)
			continue
		}
		charCount[char]++
		// 移动窗口左边指针
		for charCount[char] > 2 {
			leftChar := s[left]
			charCount[leftChar]--
			left++
		}

		// 每移动一次则计算一次
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
func main() {
	s := "ABC123"
	var c byte = 'B'
	res := longestCoincidentSubString(s, c)
	fmt.Println(res)
}
