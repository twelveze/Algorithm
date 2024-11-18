package main

import (
	"fmt"
	"sort"
)

func getKth(lo, hi, k int) int {
	sortSlice := make([][]int, 0)
	weight := make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		weight[i-lo] = getWeight(i)
	}
	for i := lo; i <= hi; i++ {
		var temp []int
		temp = append(temp, i)
		temp = append(temp, weight[i-lo])
		sortSlice = append(sortSlice, temp)
	}
	sort.Slice(sortSlice, func(i, j int) bool {
		if sortSlice[i][1] == sortSlice[j][1] {
			return sortSlice[i][0] < sortSlice[j][0]
		} else {
			return sortSlice[i][1] < sortSlice[j][1]
		}
	})
	return sortSlice[k-1][0]
}
func getWeight(i int) int {
	numWeight := 0
	for i != 1 {
		if i%2 == 1 {
			i = i*3 + 1
		} else {
			i = i / 2
		}
		numWeight++
	}
	return numWeight
}

func main() {
	lo := 12
	hi := 15
	k := 2
	res := getKth(lo, hi, k)
	fmt.Println(res)
}
