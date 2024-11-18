package main

import "fmt"

func mySqrt(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	left := 1
	right := n
	for left < right {
		mid := (left+right)/2 + 1 //思考这里为啥要+1？为什么不+1就不行
		if mid*mid > n {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
func main() {
	n := 5
	res := mySqrt(n)
	fmt.Println(res)
}
