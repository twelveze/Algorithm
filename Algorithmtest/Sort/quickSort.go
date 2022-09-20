package main

import "fmt"

func quickSort(nums []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(nums, left, right)
		quickSort(nums, left, partitionIndex-1)
		quickSort(nums, partitionIndex+1, right)
	}
	return nums
}
func partition(nums []int, left, right int) int { //这一部分是快排算法的重中之重，熟记
	pivot := left      //基准
	index := pivot + 1 //从基准右边开始遍历
	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}
func main() {
	nums := []int{1, 4, -2, 3, 9}
	nums = quickSort(nums, 0, len(nums)-1)
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
}
