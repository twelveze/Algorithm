package main

import (
	"Algorithm/Algorithmtest/LeetCodeDaily/common"
	"fmt"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	var resInterval [][]int
	if len(intervals) == 0 {
		return [][]int{}
	}
	merge := false
	left := newInterval[0]
	right := newInterval[1]
	for _, interval := range intervals {
		if interval[0] > right {
			if !merge {
				resInterval = append(resInterval, []int{left, right})
				merge = true
			}
			resInterval = append(resInterval, interval)
		} else if interval[1] < left {
			resInterval = append(resInterval, interval)
		} else {
			left = common.Min(left, interval[0])
			right = common.Max(right, interval[1])
		}
	}
	if !merge {
		resInterval = append(resInterval, []int{left, right})
	}
	return resInterval
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	res := insert(intervals, newInterval)
	fmt.Printf("%v", res)
}
