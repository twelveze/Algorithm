package main

import (
	"fmt"
	"math/rand"
)

type SolutionFor398  map[int][]int


func ConstructorFor398(nums []int) SolutionFor398 {
	position := map[int][]int{}
	for i, v := range nums{
		position[v] = append(position[v], i)
	}
	return position
}


func (pos *SolutionFor398) Pick(target int) int {
	position := *pos
	index := position[target]
	return index[rand.Intn(len(index))]
}

func main() {
	nums := []int{1,2,3,3,3}
	solution := ConstructorFor398(nums)
	res := solution.Pick(3)
	fmt.Println(res)
}