package main

import (
	"fmt"
	"math"
)

func maxRotateFunction1(nums []int) int{ //暴力超时
	len := len(nums)
	res := math.MinInt32
	for i := 0; i < len; i++{
		tempNums := nums
		nums = append(tempNums[1:], tempNums[0])
		temp := getSumForNum(nums)
		if temp > res{
			res = temp
		}
	}
	return res
}
func getSumForNum(nums []int) int{
	sum := 0
	for i := range nums{
		sum += i * nums[i]
	}
	return sum
}

//F(0) = 0 * nums[0] + 1 * nums[1] + ... + (n - 1) * nums[n - 1]
//F(1) = 1 * nums[0] + 2 * nums[1] + ... + 0 * nums[n - 1] = F(0) + numSum - n * nums[n - 1]
//更一般的，当 1 < K < n时，
//F(K) = F(k - 1) + numSum - n * nums[n - k]
func maxRotateFunction2(nums []int) int{
	numSum := 0
	for i := range nums{
		numSum += nums[i]
	}
	tempSum := 0
	for i, num := range nums{
		tempSum += i * num
	}
	res := tempSum
	n := len(nums)
	for i := len(nums) - 1; i > 0; i-- {
		tempSum += numSum - n * nums[i]
		res = max1(res, tempSum)
	}
	return res
}
func max1(a, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}
func main() {
	nums := []int{4,3,2,6}
	res := maxRotateFunction1(nums)
	fmt.Println(res)
}
