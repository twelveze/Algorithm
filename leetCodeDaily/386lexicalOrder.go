package main

import "fmt"

func lexicalOrder(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	res := make([]int, n)
	tempNum := 1
	for i := range res {
		res[i] = tempNum
		if tempNum*10 <= n { //1 10 100...
			tempNum *= 10
		} else {
			for tempNum%10 == 9 || tempNum+1 > n { //返回上一级
				tempNum /= 10
			}
			tempNum++
		}
	}
	return res
}

func main() {
	n := 13
	res := lexicalOrder(n)
	fmt.Println(res)
}
