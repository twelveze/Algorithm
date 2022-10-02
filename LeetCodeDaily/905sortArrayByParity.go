package main

import "fmt"

func sortArrayByParity(nums []int) []int{
	left, right := 0, len(nums) - 1
	for left < right{
		for left < right && nums[left] % 2 == 0{
			left++
		}
		for left < right && nums[right] % 2 == 1{
			right--
		}
		if left < right{
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}

func main() {
	nums := []int{0,2}
	res := sortArrayByParity(nums)
	fmt.Println(res)
}
