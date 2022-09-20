package main

import "fmt"

//https://leetcode.cn/problems/optimal-partition-of-string/
func partitionString(s string) int {
	var res = 1
	repeat := make([]bool, 26)
	for i := 0; i < len(s); i++ {
		char := s[i] - 'a'
		if repeat[char] {
			res++
			for j := 0; j < len(repeat); j++ {
				repeat[j] = false
			}
		}
		repeat[char] = true
	}
	return res
}
func main() {
	str := "abacaba"
	res := partitionString(str)
	fmt.Println(res)
}
