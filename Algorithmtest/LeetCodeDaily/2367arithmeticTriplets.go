package main

import "fmt"

func arithmeticTriplets(nums []int, diff int) int {
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					res++
				}
			}
		}
	}
	return res
}
func main() {
	arr := []int{0, 1, 4, 6, 7, 10}
	res := arithmeticTriplets(arr, 3)
	fmt.Println(res)
}
