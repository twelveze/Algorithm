package main

import "fmt"

func InsertionSort(arr []int){
	len := len(arr)
	for i := 1; i < len; i++{
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current{
			arr[preIndex + 1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex + 1] = current
	}
}
func main() {
	nums := []int{1,4,-2,3,9,-8}
	InsertionSort(nums)
	fmt.Println(nums)
}