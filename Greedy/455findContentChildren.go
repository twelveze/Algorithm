package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	for i, j := 0, 0; i < len(g) && j < len(s); i++ {
		//这里判断的顺序不能颠倒，如果先判断s[j]<g[i]会爆数组越界
		for j < len(s) && s[j] < g[i] {
			j++
		}
		if j < len(s) {
			res++
			j++
		}
	}
	return res
}
func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	res := findContentChildren(g, s)
	fmt.Println(res)
}
