package main

import "fmt"

func countSubIslands(grid1 [][]byte, grid2 [][]byte) int {
	var ans int
	var flag bool
	n := len(grid2)
	m := len(grid2[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid2[i][j] == '1' {
				flag = true
				//布尔类型一定要传地址！
				dfsForSubIsland(grid1, grid2, i, j, &flag)
				if flag {
					ans++
				}
			}
		}
	}
	return ans
}
func dfsForSubIsland(grid1, grid2 [][]byte, i, j int, flag *bool) {
	if i < 0 || i >= len(grid2) || j < 0 || j >= len(grid2[0]) {
		return
	}
	if grid1[i][j] == '0' && grid2[i][j] == '1' {
		*flag = false
	}
	if grid2[i][j] == '1' {
		grid2[i][j] = '0'
		dfsForSubIsland(grid1, grid2, i-1, j, flag)
		dfsForSubIsland(grid1, grid2, i+1, j, flag)
		dfsForSubIsland(grid1, grid2, i, j-1, flag)
		dfsForSubIsland(grid1, grid2, i, j+1, flag)

	}
}
func main() {
	var island1 = [][]byte{
		{'1', '1', '1', '0', '0'},
		{'0', '1', '1', '1', '1'},
		{'0', '0', '0', '0', '0'},
		{'1', '0', '0', '0', '0'},
		{'1', '1', '0', '1', '1'}}
	var island2 = [][]byte{
		{'1', '1', '1', '0', '0'},
		{'0', '0', '1', '1', '1'},
		{'0', '1', '0', '0', '0'},
		{'1', '0', '1', '1', '0'},
		{'0', '1', '0', '1', '0'}}
	res := countSubIslands(island1, island2)
	fmt.Println(res)
}
