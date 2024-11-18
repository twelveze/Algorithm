package main

import "fmt"

func maxSum(grid [][]int) int {
	n := len(grid)    //矩阵的长
	m := len(grid[0]) //矩阵的宽
	maxSum := 0
	//首先列举出所有沙漏所在3*3矩阵的左上角位置,然后遍历这个矩阵
	for i := 0; i < n-2; i++ {
		for j := 0; j < m-2; j++ {
			var tempSum int
			for a := i; a < i+3; a++ {
				for b := j; b < j+3; b++ {
					//跳过沙漏的中间位置
					if a == i+1 && b == j {
						continue
					}
					if a == i+1 && b == j+2 {
						continue
					}
					tempSum += grid[a][b]
				}
			}
			if tempSum > maxSum {
				maxSum = tempSum
			}
		}
	}
	return maxSum
}
func main() {
	grid := [][]int{{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9}}
	res := maxSum(grid)
	fmt.Println(res)
}
