package main

import "fmt"

func shortestToChar(s string, c byte) []int {
	cIndex := make([]int, 0)
	for i, char := range s {
		if char == int32(c) {
			cIndex = append(cIndex, i)
		}
	}
	res := make([]int, len(s))
	for i := range s {
		min := len(s)
		for j := 0; j < len(cIndex); j++ {
			if ABS(i, cIndex[j]) < min {
				min = ABS(i, cIndex[j])
			}
		}
		res[i] = min
	}
	return res
}
func ABS(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	s := "loveleetcode"
	var c byte = 'e'
	res := shortestToChar(s, c)
	fmt.Println(res)
}
