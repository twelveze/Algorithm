package main

import (
	"Algorithm/Model"
	"fmt"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][1] < jobs[j][1] })

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		//Search函数采用二分法搜索找到[0, i)区间内最小的满足f(i)==true的值i
		//此代码看的还是有点懵,与题解思想不符
		k := sort.Search(i, func(k int) bool { return jobs[k][1] > jobs[i-1][0] })
		dp[i] = Model.Max(dp[i-1], dp[k]+jobs[i-1][2])
	}
	return dp[n]
}

func main() {
	startTime := []int{1, 2, 3, 3}
	endTime := []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}
	res := jobScheduling(startTime, endTime, profit)
	fmt.Println(res)
}
