package main

import (
	"Algorithm/Model"
	"fmt"
)

func lengthOfLongestNoRepeatSubstring(str string) int {
	index := [128]int{}
	for i, _ := range index {
		index[i] = 0
	}
	if len(str) == 0 {
		return 0
	}
	ans := 0
	for start, end := 0, 0; end < len(str); end++ {
		if index[str[end]] != 0 {
			start = Model.Max(start, index[str[end]])
		}
		if (end - start + 1) > ans {
			ans = end - start + 1
		}
		index[str[end]] = end + 1
	}
	return ans
}
func main() {
	str := "abba"
	res := lengthOfLongestNoRepeatSubstring(str)
	fmt.Println(res)
}
