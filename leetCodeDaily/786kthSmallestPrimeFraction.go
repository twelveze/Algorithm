package main

import (
	"container/heap"
	"fmt"
	//"sort"
)

type pair struct {
	x, y int
}

//func kthSmallestPrimeFraction(arr []int, k int) []int{ //暴力
//	frac := make([]pair, 0)
//	for i, x := range arr{
//		for _, y := range arr[i + 1:]{
//			frac = append(frac, pair{x, y})
//		}
//	}
//	sort.Slice(frac, func(i, j int) bool {
//		a, b := frac[i], frac[j]
//		return a.x * b.y < a.y * b.x //这里实际上就是分数比大小的乘积形式
//	})
//	return []int{frac[k - 1].x, frac[k - 1].y}
//}

func kthSmallestPrimeFraction(arr []int, k int) []int { //使用优先队列(这里对应最小堆)
	n := len(arr)
	hp := make(hp, 0)
	for j := 1; j < n; j++ {
		hp = append(hp, frac{arr[0], arr[j], 0, j})
	}
	heap.Init(&hp)
	for loop := k - 1; loop > 0; loop-- {
		f := heap.Pop(&hp).(frac)
		if f.i+1 < f.j {
			heap.Push(&hp, frac{arr[f.i+1], f.y, f.i + 1, f.j})
		}
	}
	return []int{hp[0].x, hp[0].y}
}

type frac struct{ x, y, i, j int }
type hp []frac

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Less(i, j int) bool {
	return h[i].x*h[j].y < h[i].y*h[j].x
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(frac))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func main() {
	arr := []int{1, 2, 3, 5}
	res := kthSmallestPrimeFraction(arr, 3)
	fmt.Println(res)
}
