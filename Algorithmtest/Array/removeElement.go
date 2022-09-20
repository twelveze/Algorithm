package main

import "fmt"

func removeElement(nums []int, val int) int{
	for i := 0; i < len(nums);{
		if nums[i] == val{
			nums = append(nums[:i],nums[i+1:]...)//相当于把第i个值删除了，把第i个值前面的和第i个值后面的拼接起来
		}else{
			i++
		}
	}
	return len(nums)
}
func main() {
	nums := []int{3,2,2,3}
	res := removeElement(nums, 3)
	fmt.Println(res)
}
