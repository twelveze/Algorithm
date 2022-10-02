package main

import "fmt"

func problemOfJump(nums [][]int)(res []int){
	for i := 0; i < len(nums); i++{
		dp := make([]int, len(nums[i]))
		for j := 0; j < len(nums[i]); j++{
			index := j
			sum := 0
			for index <=len(nums[i]) - 1{
				sum += nums[i][index]
				index = index + nums[i][index]
			}
			dp[j] = sum
		}
		ans := 0
		for _, v := range dp{
			if v > ans{
				ans = v
			}
		}
		res = append(res, ans)
	}
	return
}
func main() {
	var T int
	fmt.Scan(&T)
	nums := make([][]int, T)
	for i := 0; i < T; i++{
		num := make([]int, 0)
		var n int
		fmt.Scan(&n)
		for j := 0; j < n; j++{
			var m int
			fmt.Scan(&m)
			num = append(num, m)
		}
		nums[i] = num
	}
	res := problemOfJump(nums)
	for _, v := range res{
		fmt.Println(v)
	}
}
