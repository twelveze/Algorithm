package main

import (
	"fmt"
)

func heapSort(arr []int) []int{
	len := len(arr)
	buildMaxHeap(arr, len)
	for i := len - 1; i >= 0; i--{
		swap(arr, i, 0)
		len -= 1
		heapify(arr, 0, len)
	}
	return arr
}
func buildMaxHeap(arr []int, arrLen int){
	for i := arrLen / 2; i >= 0; i--{
		heapify(arr, i, arrLen)
	}
}
func heapify(arr []int, i, arrLen int){
	left := i * 2 + 1
	right := i * 2 + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest]{
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest]{
		largest = right
	}
	if largest != i{
		swap(arr, i ,largest)
		heapify(arr, largest, arrLen)
	}
}
func swap(arr []int, a, b int){
	arr[a], arr[b] = arr[b], arr[a]
}
func main() {
	nums := []int{60,96,13,35,65,46,65,10,30,20,31,77,87,22}
	res := heapSort(nums)
	fmt.Println(res)
}
