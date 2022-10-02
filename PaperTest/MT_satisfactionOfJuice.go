package main

import (
	"fmt"
)

func satisfactionOfJuice(tasty []int) bool{
	sum := 0
	for _, v := range tasty{
		if v < 0{
			return true
		}
		sum += v
	}
	tastySum := make([][]int, len(tasty))
	for i := 0; i < len(tasty); i++{
		tastySum[i] = make([]int, len(tasty))
	}
	for i, v := range tasty{ //计算出 1～i的饮料相加得到的口味度
		tastySum[0][i] += v
		if tastySum[0][i] > sum{ //如果满足直接返回就好了
			return true
		}
	}
	for i := 1; i < len(tasty); i++{ //开始计算 l～n的饮料相加的满意度
		for j := i; j < len(tasty); j++{
			tastySum[i][j] = tastySum[i - 1][j] - tasty[i - 1]
			if tastySum[i][j] > sum{
				return true
			}
		}
	}
	return false
}

func main() {
	tasty := []int{-5,0,5}
	res := satisfactionOfJuice(tasty)
	fmt.Println(res)
}
