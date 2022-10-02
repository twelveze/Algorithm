package main

import "fmt"

func firstUniqChar(s string) int{
	index := make([]int, 26)
	for i, c := range s{
		index[c - 'a'] = i
	}
	for i, c := range s{
		if index[c - 'a'] == i{
			return i
		} else{
			return -1
		}
	}
	return -1
}
func main() {
	s := "hello"
	res := firstUniqChar(s)
	fmt.Println(res)
}
