package main

import "fmt"

// leetcode - 63 不同路径
func uniquePathsWithObstacles(Grid [][]int) int {
	n := len(Grid)
	m := len(Grid[0])

	// 初始化
	stepGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		stepGrid[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		if Grid[i][0] == 0 {
			stepGrid[i][0] = 1
		} else {
			break
		}
	}

	for j := 0; j < m; j++ {
		if Grid[0][j] == 0 {
			stepGrid[0][j] = 1
		} else {
			break
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if Grid[i][j] == 0 {
				stepGrid[i][j] = stepGrid[i-1][j] + stepGrid[i][j-1]
			} else {
				stepGrid[i][j] = 0
			}
		}
	}
	return stepGrid[n-1][m-1]
}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
