package main

import "fmt"

func buildArray(target []int, n int) []string {
	var ans []string
	index := 0
	for _, v := range target {
		//这里就是中间隔了多少个没有，就插入多少个push pop对
		for i := 0; i < v-index-1; i++ {
			ans = append(ans, "PUSH", "POP")
		}
		index = v
		ans = append(ans, "PUSH")
	}
	return ans
}
func main() {
	target := []int{1, 2, 3}
	n := 3
	res := buildArray(target, n)
	fmt.Println(res)
}
