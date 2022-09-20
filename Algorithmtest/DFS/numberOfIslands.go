package main

import "fmt"

func numberOfIsland(islands [][]byte) int {
	var nums int
	visit := make([][]bool, len(islands))
	for i := range visit {
		visit[i] = make([]bool, len(islands[0]))
		for j := 0; j < len(visit[i]); j++ {
			visit[i][j] = false
		}
	}
	for i := 0; i < len(islands); i++ {
		for j := 0; j < len(islands[i]); j++ {
			if !visit[i][j] && islands[i][j] == '1' {
				nums++
				visit[i][j] = true
				dfsForIsland(islands, visit, i, j)
			}
		}
	}
	return nums
}
func dfsForIsland(island [][]byte, visit [][]bool, i, j int) {
	if i < 0 || i > len(visit) || j < 0 || j > len(visit[0]) {
		return
	}
	if i-1 >= 0 && !visit[i-1][j] && island[i-1][j] == '1' {
		visit[i-1][j] = true
		dfsForIsland(island, visit, i-1, j)
	}
	if i+1 < len(visit) && !visit[i+1][j] && island[i+1][j] == '1' {
		visit[i+1][j] = true
		dfsForIsland(island, visit, i+1, j)
	}
	if j-1 >= 0 && !visit[i][j-1] && island[i][j-1] == '1' {
		visit[i][j-1] = true
		dfsForIsland(island, visit, i, j-1)
	}
	if j+1 < len(visit[0]) && !visit[i][j+1] && island[i][j+1] == '1' {
		visit[i][j+1] = true
		dfsForIsland(island, visit, i, j+1)
	}
}
func main() {
	var islands = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	res := numberOfIsland(islands)
	fmt.Println(res)
}
