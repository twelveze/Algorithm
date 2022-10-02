package main

import (
	"Algorithm/Algorithmtest/Model"
	"fmt"
)

const MAX = 1000 + 50
const MOD = 1e9 + 7

//https://leetcode.cn/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/solution/by-tsreaper-qzp3/
//组合数思想,k步里有delta步是正向走，k-delta步反向走，求出组合的组合数
func numberOfWays(startPos, endPos, k int) int {
	comb := make([][]int, MAX)
	for i := 0; i < MAX; i++ {
		comb[i] = make([]int, MAX)
	}
	//abs为向目标方向移动的直线距离，当k > asb时需要先往反方向移动再朝目标方向移动
	absRes := Model.Abs(startPos, endPos)
	if (k-absRes)%2 == 1 || absRes > k {
		return 0
	}
	//需要往目标方向移动的距离
	delta := absRes + (k-absRes)/2
	comb[0][0] = 1
	for i := 0; i <= k; i++ {
		comb[i][0] = 1 //初始化
		comb[i][i] = 1 //初始化
		for j := 1; j < i; j++ {
			//求组合数，实际上是数学概念的具体化
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % MOD
		}
	}
	return comb[k][delta]
}

func main() {
	startPos := 1
	endPos := 2
	k := 3
	res := numberOfWays(startPos, endPos, k)
	fmt.Println(res)
}
