package main

import "fmt"

func ncov_defect(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	var res int
	var visit = make([][]int, m)
	for i := range visit {
		visit[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i-1 >= 0 && grid[i-1][j] == 0 && visit[i-1][j] == 0 {
					res++
					visit[i-1][j] = 1
				}
				if j-1 >= 0 && grid[i][j-1] == 0 && visit[i][j-1] == 0 {
					res++
					visit[i][j-1] = 1
				}
				if i+1 < m && grid[i+1][j] == 0 && visit[i+1][j] == 0 {
					res++
					visit[i+1][j] = 1
				}
				if j+1 < n && grid[i][j+1] == 0 && visit[i][j+1] == 0 {
					res++
					visit[i][j+1] = 1
				}
			}
		}
	}
	return res
}
func main() {
	var grid = [][]int{{0, 1, 0, 0, 0, 0, 0, 1}, {0, 1, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 0}}
	res := ncov_defect(grid)
	fmt.Println(res)
}
