package main

import "fmt"

func rotate(nums []int, k int){
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}
func reverse(arr []int){
	for i := 0; i < len(arr)/2; i++{
		temp := arr[i]
		arr[i] = arr[len(arr) - 1 - i]
		arr[len(arr) - 1 - i] = temp
	}
}
func main() {
	nums := []int{1,2,3,4,5,6,7}
	rotate(nums, 3)
	fmt.Println(nums)
}