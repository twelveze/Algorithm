package main

import (
	"fmt"
	"unicode"
)

func letterCasePermutation(s string) []string {
	var dfs func([]byte, int)
	var ans []string
	dfs = func(str []byte, pos int) {
		for pos < len(str) && unicode.IsDigit(rune(str[pos])) {
			pos++
		}
		if pos == len(str) {
			ans = append(ans, string(str))
			return
		}
		dfs(str, pos+1)
		str[pos] ^= 32
		dfs(str, pos+1)
	}
	dfs([]byte(s), 0)
	return ans
}

func main() {
	s := "a1b2"
	res := letterCasePermutation(s)
	fmt.Println(res)
}
