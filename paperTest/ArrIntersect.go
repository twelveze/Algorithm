package main

import (
	"fmt"
	"sort"
)

func ArrIntersect(a, b []int) (res []int) {
	lenA := len(a)
	lenB := len(b)
	indexA, indexB := 0, 0
	sort.Ints(a)
	sort.Ints(b)
	for indexA < lenA && indexB < lenB {
		if a[indexA] < b[indexB] {
			indexA++
		} else if a[indexA] > b[indexB] {
			indexB++
		} else {
			res = append(res, a[indexA])
			indexA++
			indexB++
		}
	}
	return
}
func main() {
	a := []int{4, 4, 5, 3, 1}
	b := []int{2, 3, 4, 4, 7}
	res := ArrIntersect(a, b)
	fmt.Println(res)
}
