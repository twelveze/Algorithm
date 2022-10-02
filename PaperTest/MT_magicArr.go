package main

import "fmt"

func magicArr(arr []int) int{
	max := arr[0]
	sum := 0
	index := 0
	for i := 1; i < len(arr); i++{
		if arr[i] + arr[i - 1] > arr[i]{
			arr[i] += arr[i - 1]
		}
		if arr[i] > max{
			index = i
			max = arr[i]
		}
	}
	sum += max * 10
	max = arr[index + 1]
	for i := index + 1; i < len(arr); i++{
		if arr[i] + arr[i - 1] > arr[i]{
			arr[i] += arr[i - 1]
		}
		if arr[i] > max{
			index = i
			max = arr[i]
		}
	}
	sum += max * 10
	return sum
}
func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++{
		var num int
		fmt.Scan(&num)
		arr[i] = num
	}
	res := magicArr(arr)
	fmt.Println(res)
}
