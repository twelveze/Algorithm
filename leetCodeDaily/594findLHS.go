package main

import "fmt"

func findHS(nums []int) (ans int) {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if c1 := cnt[num+1]; c1 > 0 && c+c1 > ans {
			ans = c + c1
		}
	}
	return
}
func main() {
	var nums = []int{1, 3, 2, 2, 5, 2, 3, 7}
	res := findHS(nums)
	fmt.Println(res)
}
