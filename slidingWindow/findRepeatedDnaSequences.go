package main

import "fmt"

// leetcode - 187 重复的DNA序列

func findRepeatedDnaSequences(s string) []string {
	dnaSequenceMap := make(map[string]int, 0)
	ans := make([]string, 0)
	for i := 0; i < len(s)-9; i++ {
		subDNA := s[i : i+10]
		dnaSequenceMap[subDNA]++
		if count, ok := dnaSequenceMap[subDNA]; ok && count == 2 {
			ans = append(ans, subDNA)
		}
	}
	return ans
}
func main() {
	dnaStr := "AAAAAAAAAAA"
	res := findRepeatedDnaSequences(dnaStr)
	fmt.Println(res)
}
