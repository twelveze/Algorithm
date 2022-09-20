package main

import "fmt"

func rob(house []int) int{
	if len(house) == 0{
		return 0
	}
	if len(house) == 1{
		return house[0]
	}
	dp := make([]int, len(house))
	dp[0] = house[0]
	dp[1] = maxNum(house[0], house[1])
	for i := 2; i < len(house); i++{
		dp[i] = maxNum(dp[i - 2] + house[i], dp[i - 1])
	}
	return dp[len(house) - 1]
}
func maxNum(a, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}
func main() {
	house := []int{2,7,9,3,1}
	res := rob(house)
	fmt.Println(res)
}
