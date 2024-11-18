package main

import "fmt"

// 华为OD - 补种未成活胡杨
/*
	近年来，我国防沙举措取得良好成果
	某沙漠新种植N颗胡杨(编号1-N)，排成一排，一个月后，有M颗胡杨未能成活
	现可补种胡杨K颗，请问如何补种(只能补种，不能新种)，可以得到最多的连续胡杨？

	输入描述：
	N 总种植数量
	1 <= N <= 1000

	M 未成活数量
	1 <= M <= N

	未成活的数下标数组(从1开始)
	[x, x, x, ...]

	K 补种数量
	0 <= K <= M
*/

func reLiveTree(N, M int, badTree []int, K int) int {
	badTreeMap := make(map[int]struct{}, 0)
	for _, badNum := range badTree {
		badTreeMap[badNum] = struct{}{}
	}
	if M == K {
		return N
	}
	ans := 0
	badIndex := 0
	for left, right := 1, 1; right <= N; right++ {
		if _, ok := badTreeMap[right]; ok {
			if K > 0 {
				K--
			} else {
				left = badTree[badIndex] + 1
				badIndex++
			}
		}
		// 更新最新长度
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}

func main() {
	N := 10
	M := 3
	badTree := []int{2, 4, 7}
	K := 1
	res := reLiveTree(N, M, badTree, K)
	fmt.Println(res)
}
