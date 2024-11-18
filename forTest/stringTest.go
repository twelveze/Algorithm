package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	num1 := make([]int, len(nums))
	copy(num1, nums)
	fmt.Println(num1)
}
