package main

import "fmt"

func twoSum(nums []int, target int) []int{
	res := make([]int, 0)
	m := make(map[int]int, 0)
	for i, v := range nums{
		if val,exist := m[target - v];exist{
			res = append(res, val)
			res = append(res, i)
		}
		m[v] = i
	}
	return res
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
