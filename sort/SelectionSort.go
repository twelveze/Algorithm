package main

import "fmt"

func SelectionSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
func main() {
	nums := []int{1, 4, -2, 3, 9, -8}
	SelectionSort(nums)
	fmt.Println(nums)
}
