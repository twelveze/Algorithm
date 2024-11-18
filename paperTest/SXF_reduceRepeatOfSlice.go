package main

import (
	"fmt"
	"sort"
)

func unique(s []int) int {
	sort.Ints(s)
	index := 1
	for index <= len(s)-1 {
		if s[index] == s[index-1] {
			temp := s[index]
			s = append(s[:index], s[index+1:]...)
			s = append(s, temp)
		} else {
			index++
			if s[index] < s[index-1] {
				break
			}
		}
	}
	return index
}
func main() {
	s := []int{3, 2, 3, 1, 5, 2, 1, 1}
	res := unique(s)
	fmt.Println(res)
}
