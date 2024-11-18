package main

import (
	"fmt"
	"math"
	"model"
)

//https://leetcode.cn/problems/minimum-falling-path-sum/
func minFallingPathSum(matrix [][]int) int {
	var memo [][]int
	n := len(matrix)
	//初始化备忘录
	for i := 0; i < n; i++ {
		temp := make([]int, n)
		for j := 0; j < len(temp); j++ {
			temp[j] = 10001
		}
		memo = append(memo, temp)
	}

	res := math.MaxInt
	for i := 0; i < len(matrix[n-1]); i++ {
		res = model.MIN(res, dp(memo, matrix, n-1, i))
	}
	return res
}
func dp(memo, matrix [][]int, n, i int) int {
	if n < 0 || i < 0 || n >= len(matrix) || i >= len(matrix[0]) {
		return 10000 //这里大于等于9900就行
	}
	if n == 0 {
		return matrix[0][i]
	}
	//查找备忘录，防止重复计算
	if memo[n][i] != 10001 {
		return memo[n][i]
	}
	memo[n][i] = matrix[n][i] + tripleMin(dp(memo, matrix, n-1, i), dp(memo, matrix, n-1, i-1), dp(memo, matrix, n-1, i+1))
	return memo[n][i]
}
func tripleMin(a, b, c int) int {
	return model.MIN(a, model.MIN(b, c))
}
func main() {
	row1 := []int{-19, 57}
	row2 := []int{-40, -5}
	var matrix [][]int
	matrix = append(matrix, row1, row2)
	res := minFallingPathSum(matrix)
	fmt.Println(res)
}
