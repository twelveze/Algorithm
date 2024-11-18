package main

import (
	"fmt"
	"math"
)

//func getLeftMaxOrRightMax(temp int, flag int, end int, n int) int{
//	if flag + 1 == end || (end != n && flag + 2 == end){
//		return temp
//	}else{
//		nextTemp := ( flag + 1 + end ) / 2
//		temp = getLeftMaxOrRightMax(temp + nextTemp, nextTemp, end, n)
//		return temp
//	}
//}
//func getTempRes(i, n int) int{
//	leftMax, rightMax := i, i
//	leftTemp := ( 1 + i - 1 ) / 2
//	rightTemp := (i + 1 + n) / 2
//	if leftTemp > 1 || leftMax - 2 == 1{
//		leftMax = getLeftMaxOrRightMax(i + leftTemp, leftTemp, i, n)
//	}
//	if rightTemp < n{
//		rightMax = getLeftMaxOrRightMax(i + rightTemp, rightTemp, n, n)
//	}
//	if leftMax > rightMax{
//		return leftMax
//	}else{
//		return rightMax
//	}
//}
//func getMoneyAmount(n int) int{
//	resArr := make([]int, n + 1)
//	if n == 1{
//		return 0
//	}
//	for i := 1; i <= n; i++{
//		resArr[i] = getTempRes(i, n)
//	}
//	res := math.MaxInt64
//	for i := 1; i <= n; i++{
//		if resArr[i] < res{
//			res = resArr[i]
//		}
//	}
//	return res
//}
func getMoneyAmount(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			f[i][j] = minCost
		}
	}
	return f[1][n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func main() {
	var n int // 1<= n <= 200
	n = 10
	//res := getMoneyAmount(n) //错误方法的思想是对与1 - n，求出每一个数字可以完成猜想的最大money，然后取其中最小的值
	//https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/solution/cai-shu-zi-da-xiao-ii-by-leetcode-soluti-a7vg/
	res := getMoneyAmount(n)
	fmt.Println(res)
}
