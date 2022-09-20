package main

import "fmt"

//归并排序
func MergeSort(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	middle := len / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(MergeSort(left), MergeSort(right))
}
func merge(left []int, right []int) []int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		res = append(res, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		res = append(res, right[0])
		right = right[1:]
	}
	return res
}
func main() {
	nums := []int{1, 4, -2, 3, 9}
	res := MergeSort(nums)
	fmt.Println(res)
}
