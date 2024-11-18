package main

import (
	"fmt"
	"sort"
)

func howMuch(bouns []int) int {
	sort.Ints(bouns)
	res := 0
	for len(bouns) > 1 {
		if bouns[0] != bouns[len(bouns)-1] {
			temp := bouns[len(bouns)-1] - bouns[0]
			res += temp
			bouns = bouns[1 : len(bouns)-1]
			bouns = append(bouns, temp)
		} else if bouns[0] == bouns[len(bouns)-1] {
			bouns = bouns[1:]
		}
	}
	return res
}
func main() {
	var nums int
	fmt.Scanf("%d", &nums)
	bouns := make([]int, 0)
	for i := 0; i < nums; i++ {
		var temp int
		fmt.Scan(&temp)
		bouns = append(bouns, temp)
	}
	res := howMuch(bouns)
	fmt.Println(res)
}
