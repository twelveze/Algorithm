package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, i := range nums1 {
		m[i] += 1
	}
	k := 0
	for _, j := range nums2 {
		if m[j] != 0 {
			m[j] -= 1
			m[k] = j
			k++
		}
	}
	return nums2
}
func main() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{2, 3, 4}
	res := intersect(nums1, nums2)
	fmt.Println(res)
}
