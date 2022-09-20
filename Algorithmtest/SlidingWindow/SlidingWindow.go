package main

import "fmt"

func SlidingWindow(nums []int, k int) []int{
	if len(nums) == 0{
		return []int{}
	}
	var queue []int
	var result []int
	for i,_ := range nums{
		for len(queue) > 0 && nums[i] > queue[len(queue) - 1]{
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, nums[i])
		if i >= k && nums[i - k] == queue[0]{
			queue = queue[1:]
		}
		if i >= k - 1{
			result = append(result, queue[0])
		}
	}
	return result
}
func main() {
	nums := []int{5,3,-1,-3,5,3,4,7}
	queue := SlidingWindow(nums, 3)
	fmt.Println(queue)
}
