package main

import (
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
		//二分查找找到结束时间小于等于第i-1份工作的工作下标
		k := sort.Search(i, func(j int) bool { return jobs[j][1] > jobs[i-1][0] })
		dp[i] = max(dp[i-1], dp[k]+jobs[i-1][2])
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
