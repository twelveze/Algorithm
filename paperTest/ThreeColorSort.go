package main

import (
	"fmt"
)

func ThreeColorSort(arr []int) []int {
	index0 := 0
	index2 := len(arr) - 1
	for i := 0; i < index2; i++ {
		if arr[i] == 0 {
			arr[i], arr[index0] = arr[index0], arr[i]
			index0++
		} else if arr[i] == 2 {
			arr[i], arr[index2] = arr[index2], arr[i]
			index2--
		} else if arr[i] == 1 {
			continue
		}
	}
	return arr
}
func main() {
	arr := make([]int, 0)
	var value int
	for i := 0; i < 10; i++ {
		_, err := fmt.Scanf("%d", &value)
		if err != nil {
		}
		arr = append(arr, value)
	}
	res := ThreeColorSort(arr)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d ", res[i])
	}
}
