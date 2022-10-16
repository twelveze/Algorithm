package main

import "fmt"

//深度优先搜索
//用两种颜色对图进行染色，如果可以完成染色，那么就说明可以将所有人分进两组
//初始化所有人的颜色为0，表示还没有染色
//遍历所有人，如果当前人没有染色，那么就用颜色1染色
//然后将其所有不喜欢的人用颜色2染色，如果染色过程中出现了冲突，那么就说明无法将所有人分进两组
func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, e := range dislikes {
		a, b := e[0]-1, e[1]-1
		g[a] = append(g[a], b) //双向的dislike
		g[b] = append(g[b], a)
	}

	color := make([]int, n)
	var dfs func(int, int) bool

	dfs = func(i int, c int) bool {
		color[i] = c
		for _, j := range g[i] {
			if color[j] == c { //出现了冲突，直接返回false
				return false
			}
			if color[j] == 0 && !dfs(j, 3^c) { //只有没遍历过的才进行遍历，1的dislike变2，2的dislike变1
				return false
			}
		}
		return true
	}
	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}
func main() {
	dislikes := [][]int{{1, 2}, {1, 3}, {2, 4}}
	n := 4
	res := possibleBipartition(n, dislikes)
	fmt.Println(res)
}
