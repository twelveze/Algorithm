package main

import "fmt"

func chooseMonkeyKing(M int, N int) int{
	monkey := make([]int, N)
	for i := 0; i < N; i++{
		monkey[i] = i + 1
	}
	index := 0
	i := 1
	remainMonkey := N
	for remainMonkey > 1{
		if monkey[index] != -1{
			if i == M{
				monkey[index] = -1
				remainMonkey--
				M++
			}
			i++
			if i > M{
				i  = 1
			}
			index++
			if index >= remainMonkey{
				index = index % remainMonkey
			}
		}
	}
	var res int
	for i := 0; i < N; i++{
		if monkey[i] != -1{
			res = monkey[i] + 1
			if res > N{
				res = res % N
			}
		}
	}
	return res % 2
}
func main() {
	M := 3
	N := 5
	res := chooseMonkeyKing(M, N)
	fmt.Println(res)
}
