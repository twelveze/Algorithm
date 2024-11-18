package main

import "fmt"

func findAnagrams(s, p string) (ans []int) {
	if len(p) > len(s) {
		return
	}
	var sCount, pCount [26]int
	for i, c := range p {
		sCount[s[i]-'a']++
		pCount[c-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}
	for i := range s[:len(s)-len(p)] {
		sCount[s[i]-'a']--
		sCount[s[i+len(p)]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}
