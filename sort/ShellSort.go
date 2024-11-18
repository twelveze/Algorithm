package main

import "fmt"

func ShellSort(arr []int) { //建议用3x + 1作为起始分组
	length := len(arr)
	gap := 1
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
}
func main() {
	nums := []int{1, 4, -2, 3, 9}
	ShellSort(nums)
	fmt.Println(nums)
}
