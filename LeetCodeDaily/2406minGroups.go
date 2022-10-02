package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//小根堆
type smallhp struct {
	sort.IntSlice
}

func (h *smallhp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *smallhp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func minGroups(intervals [][]int) int {
	h := smallhp{}
	//先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		//如果目前的小根堆长度是0，或者待加入的分组左端点大于最小的右端点，那么可以加入同一组
		if h.Len() == 0 || intervals[i][0] > h.IntSlice[0] {
			heap.Push(&h, intervals[i][1])
		} else { //如果不能，必须重开一组
			h.IntSlice[0] = intervals[i][1]
			//调整小根堆
			heap.Fix(&h, 0)
		}
	}
	return h.Len()
}

func main() {
	intervals := [][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}
	res := minGroups(intervals)
	fmt.Println(res)
}
