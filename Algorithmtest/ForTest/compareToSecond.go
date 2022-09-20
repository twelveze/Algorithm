package main

import "fmt"

func main() {
	var arr = []int{2,3,1,6,8}
	var first int
	var second int
	for i := 0; i < len(arr); i++{
		if arr[i] > first{
			second = first
			first = arr[i]
		}else if arr[i] <= first && arr[i] > second{
			second = arr[i]
		}
	}
	fmt.Println(first, second)
}
