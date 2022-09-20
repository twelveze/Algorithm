package main

import "fmt"

func BubbleSort(arr []int){
	len := len(arr)
	for i := len - 1; i > 0; i--{
		for j := 1; j <= i; j++{
			if arr[j - 1] > arr[j]{
				arr[j - 1], arr[j] = arr[j], arr[j - 1]
			}
		}
	}
}
func main() {
	nums := []int{1,4,-2,3,9,-8}
	BubbleSort(nums)
	fmt.Println(nums)
}
