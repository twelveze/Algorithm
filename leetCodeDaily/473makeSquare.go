package main

import (
	"fmt"
	"sort"
)

//首先计算所有火柴的总长度totalLen,如果不是4的倍数，那么不可能拼成正方形
//当totalLen是4的倍数时,每个边的边长为totalLen / 4
//用edge来记录4条边已经放入的火柴总长度,使火柴放入后每个边长不超过 totalLen / 4
//当edge满足后继续看edge + 1, 如果所有的火柴都放完了,那么说明可以拼成正方形
func makeSquare(matchsticks []int) bool {
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}

	var edges = [4]int{} //初始长度为0的四个边
	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == len(matchsticks) { //如果所有的边都选完了,就可以组成一个正方形
			return true
		}
		for i := range edges {
			edges[i] += matchsticks[index]
			if edges[i] <= totalLen/4 && dfs(index+1) { //判断边长超没超,没超就继续加
				return true
			}
			edges[i] -= matchsticks[index] //边长超了就回溯，加到下一条边上
		}
		return false
	}
	return dfs(0)
}
func main() {
	matchsticks := []int{3, 3, 3, 3, 4}
	res := makeSquare(matchsticks)
	fmt.Println(res)
}
