package main

import (
	"fmt"
	"sort"
)

func threeSum(num []int) [][]int {
	res := make([][]int, 0)
	if len(num) < 0 {
		return res
	}
	sort.Ints(num)
	if num[0] > 0 {
		return res
	}
	for i := 0; i < len(num); i++ {
		left := i + 1
		right := len(num) - 1
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		for left < right {
			if num[i]+num[left]+num[right] > 0 {
				right--
			} else if num[i]+num[left]+num[right] < 0 {
				left++
			} else {
				temp := make([]int, 0)
				temp = append(temp, num[i])
				temp = append(temp, num[left])
				temp = append(temp, num[right])
				res = append(res, temp)
				left++
				right--
				//去重
				for left < right && num[left] == num[left-1] {
					left++
				}
				for left < right && num[right] == num[right+1] {
					right--
				}
			}
		}
	}
	return res
}
func main() {
	num := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(num)
	fmt.Println(res)
}
