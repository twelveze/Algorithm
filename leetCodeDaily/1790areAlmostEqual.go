package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	len := len(s1)
	first, second := -1, -1
	for i := 0; i < len; i++ {
		if s1[i] != s2[i] {
			if first < 0 {
				first = i
			} else if second < 0 {
				second = i
			} else {
				return false
			}
		}
	}
	return first < 0 || second >= 0 && s1[first] == s2[second] && s1[second] == s2[first]
}
func main() {
	s1 := "bank"
	s2 := "kanb"
	res := areAlmostEqual(s1, s2)
	fmt.Println(res)
}
