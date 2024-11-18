package main

import (
	"fmt"
	"math"
)

func findNthDigit(n int) int {
	cur, base, max := 1, 9, math.MaxInt64
	for n > cur*base {
		n -= cur * base
		cur++
		base *= 10
		if cur*base > max {
			break
		}
	}
	num := int(math.Pow10(cur-1)) + (n-1)/cur
	index := (n - 1) % cur //对应在num中的第几位
	return num / int(math.Pow10(cur-index-1)) % 10
}
func main() {
	var n = 3
	res := findNthDigit(n)
	fmt.Println(res)
}
