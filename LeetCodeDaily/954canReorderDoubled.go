package main

import (
	"fmt"
	"sort"
)

func canReorderDoubled(arr []int) bool{
	count := make(map[int]int, 0)
	for _, v := range arr{
		count[v]++
	}
	if count[0]%2 == 1{
		return false
	}
	vals := make([]int, 0, len(count))
	for x := range count {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })
	for _, x := range vals {
		if count[2*x] < count[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		count[2*x] -= count[x]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	arr := []int{4, -2, 2, -4}
	res := canReorderDoubled(arr)
	if res == false{
		fmt.Println("False")
	}else{
		fmt.Println("True")
	}
}
