package main

import "fmt"

func findAnagrams(s, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	pm := map[byte]int{}
	sm := map[byte]int{}
	for i := 0; i < len(p); i++ {
		pm[p[i]]++
		sm[s[i]]++
	}
	start, end := 0, len(p)-1
	var res []int
	for end < len(s)-1 {
		if isMatch(sm, pm) {
			res = append(res, start)
		}
		sm[s[start]]--
		start++
		end++
		sm[s[end]]++
	}
	if isMatch(sm, pm) { //剩下最后一个窗口没有检测
		res = append(res, start)
	}
	return res
}
func isMatch(sm, pm map[byte]int) bool {
	for v := range sm {
		if sm[v] != pm[v] {
			return false
		}
	}
	return true
}
func main() {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}
