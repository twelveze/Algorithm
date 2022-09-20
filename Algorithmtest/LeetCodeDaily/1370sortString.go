package main

import "fmt"

func sortString(s string) string {
	var res string
	count := make([]int, 26)
	for i := range s {
		count[s[i]-'a']++
	}
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if count[i] > 0 {
				res += string(i + 'a')
				count[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if count[i] > 0 {
				res += string(i + 'a')
				count[i]--
			}
		}
	}
	return res
}
func main() {
	s := "aaaabbbbcccc"
	res := sortString(s)
	fmt.Println(res)
}
