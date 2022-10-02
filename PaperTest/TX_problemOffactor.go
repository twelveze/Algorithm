package main

import "fmt"

func problemOffactor(nums []int)(res []int){
	for i := 0; i < len(nums); i++{
		index := make([]int,4)
		index[0] = 1
		for j := 1; j < 4; j++{
			index[j] = index[j - 1] + nums[i]
		}

	}
	return
}
func main() {
	var T int
	nums := make([]int, T)
	fmt.Scan(&T)
	for i := 0; i < T; i++{
		var num int
		fmt.Scan(&num)
		nums[i] = num
	}
	res := problemOffactor(nums)
	for _, v := range res{
		fmt.Println(v)
	}
}
