package main

import (
	"Model"
	"fmt"
)

func integerReplacement(n int) int{
	if n == 1{
		return 0
	}
	if n % 2 == 0{
		return 1 + integerReplacement(n / 2)
	}else{
		return 2 + Model.MIN(integerReplacement(n / 2), integerReplacement(n / 2 + 1) )
	}
}

func main() {
	var n int
	n = 8
	res := integerReplacement(n)
	fmt.Println(res)
}
