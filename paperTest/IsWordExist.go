package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	count := 0
	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, visited, m, n, i, j, count) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, word string, visited [][]bool, m, n, i, j, count int) bool {
	if count == len(word) {
		return true
	}
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[count] {
		return false
	}
	if visited[i][j] {
		return false
	}
	visited[i][j] = true
	res := dfs(board, word, visited, m, n, i+1, j, count+1) ||
		dfs(board, word, visited, m, n, i-1, j, count+1) ||
		dfs(board, word, visited, m, n, i, j+1, count+1) ||
		dfs(board, word, visited, m, n, i, j-1, count+1)
	return res
}
func main() {
	board := [][]byte{{'X', 'Y', 'Z', 'E'}, {'S', 'F', 'Z', 'S'}, {'X', 'D', 'E', 'E'}}
	word := "XYZZED"
	res := exist(board, word)
	fmt.Println(res)
}
