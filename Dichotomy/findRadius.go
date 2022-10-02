package main

import (
	"fmt"
	"sort"
)

func findRadius(house, heaters []int)int{
	sort.Ints(house)
	sort.Ints(heaters)
	left := 1
	right := len(house)
	for left < right{
		mid := (left + right) / 2
		if check(mid, house, heaters){
			right = mid
		}else{
			left = mid + 1
		}
	}
	return left
}
func check(r int, houses, heaters []int) bool{ //判断以r为半径符不符合覆盖的条件

	for i := 0; i < len(houses); i++{
		housecounts := 0
		for j := 0; j < len(heaters); j++{
			if houses[i] < heaters[j] - r || houses[i] > heaters[j] + r{
				housecounts++
			}else{
				break
			}
		}
		if housecounts == len(heaters){
			return false
		}
	}
	return true
}
func main() {
	house := []int{1,2,3,4,5,6,7,8,9}
	heaters := []int{4,9}
	res := findRadius(house, heaters)
	fmt.Println(res)
}
