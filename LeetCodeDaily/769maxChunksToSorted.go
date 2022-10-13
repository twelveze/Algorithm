package main

import "fmt"

//若已遍历过的数中的最大值max与当前遍历到的下标i相等，说明可以进行一次分割，累加答案
func maxChunksToSorted(arr []int) int {
	ans := 0
	max := 0
	for i, v := range arr {
		if v > max {
			max = v
		}
		if max == i {
			ans++
		}
	}
	return ans
}
func main() {
	arr := []int{1, 0, 2, 3, 4}
	res := maxChunksToSorted(arr)
	fmt.Println(res)
}
