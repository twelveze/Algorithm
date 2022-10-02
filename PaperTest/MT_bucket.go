package main

import (
	"fmt"
)

func bucket(band []int, m int)int{
	length := len(band)
	max := 0
	total := 0
	for i := 0; i < len(band); i++{
		if band[i] > max{
			max = band[i]
		}
		total += band[i]
	}
	total += m
	left, right := 0, max + m
	for left < right{
		mid := (left + right) / 2
		if mid * length > total{
			right = mid - 1
		}else{
			left = mid
		}
	}
	return left
}
func main() {
	var n, m int
	band := make([]int, 0)
	fmt.Scanf("%d %d",&n, &m)
	for i := 0; i < n; i++{
		var len int
		fmt.Scan(&len)
		band = append(band, len)
	}
	res := bucket(band, m)
	fmt.Println(res)
}
