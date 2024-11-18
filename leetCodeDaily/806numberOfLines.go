package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	n := len(s)
	line := 1
	redundant := 0
	for i := 0; i < n; i++ {
		redundant += widths[s[i]-'a']
		if redundant > 100 {
			redundant = widths[s[i]-'a']
			line++
		}
	}
	return []int{line, redundant}
}

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"
	res := numberOfLines(widths, s)
	fmt.Println(res)
}
