package main

import (
	"Algorithm/Model"
	"fmt"
)

//滑动窗口
func totalFruit(fruits []int) int {
	count := make(map[int]int, 0) //记录对应水果种类的个数
	left := 0
	ans := 0
	for right, value := range fruits {
		count[value]++
		for len(count) > 2 {
			cnt := fruits[left]
			count[cnt]--
			if count[cnt] == 0 {
				delete(count, cnt)
			}
			left++
		}
		ans = Model.Max(ans, right-left+1)
	}
	return ans
}

func main() {
	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	res := totalFruit(fruits)
	fmt.Println(res)
}
