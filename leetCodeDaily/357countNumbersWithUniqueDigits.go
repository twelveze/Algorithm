package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	res, cur := 10, 9
	for i := 0; i < n-1; i++ {
		cur = cur * (9 - i) //第一位有1～9九种选择，第二位可选0～9中第一位没选的其中9种，第三位可选0～8中前两位没选过的其中8种，以此类推
		res += cur
	}
	return res
}

func main() {
	n := 2
	res := countNumbersWithUniqueDigits(n)
	fmt.Println(res)
}
