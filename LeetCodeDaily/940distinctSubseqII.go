package main

import "fmt"

//实际上要计算的是以s为结尾的字符串的数量 例如: a b c 以a结尾的是a, 以b结尾的是a ab,以c结尾的是ac bc adc c 所以是七个
func distinctSubseqII(s string) int {
	const mod int = 1e9 + 7
	ans := 0
	last := make([]int, 26) //表示以c为结尾的子字符串中'c'所在s中的index
	for i := range last {
		last[i] = -1
	}
	n := len(s)
	f := make([]int, n) //f[i]表示以s[i]为结尾字符串的子序列的数目
	for i := range f {
		f[i] = 1
	}
	for i, c := range s {
		for _, j := range last {
			if j != -1 {
				f[i] = (f[i] + f[j]) % mod
			}
		}
		last[c-'a'] = i
	}
	for _, i := range last {
		if i != -1 {
			ans = (ans + f[i]) % mod
		}
	}
	return ans
}
func main() {
	str := "abc"
	res := distinctSubseqII(str)
	fmt.Println(res)
}
