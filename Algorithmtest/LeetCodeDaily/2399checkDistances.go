package main

import "fmt"

func checkDistances(s string, distance []int) bool {
	index := make([]int, 26)
	for i := 0; i < len(index); i++ {
		index[i] = -1
	}
	for i := 0; i < len(s); i++ {
		if index[s[i]-'a'] != -1 {
			temp := i - index[s[i]-'a'] - 1
			if distance[s[i]-'a'] != temp {
				return false
			}
		} else {
			index[s[i]-'a'] = i
		}
	}
	return true
}
func main() {
	s := "abaccb"
	distance := []int{1, 2, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	res := checkDistances(s, distance)
	fmt.Println(res)
}
