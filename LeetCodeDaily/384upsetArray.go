package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums, original []int
}


func constructor(nums []int) Solution {
	return Solution{nums: nums, original: append([]int(nil), nums...)}
}

func (s *Solution) Reset() []int {
	copy(s.nums, s.original)
	return s.nums
}


func (s *Solution) Shuffle() []int { //Fisher-Yates洗牌算法
	n := len(s.nums)
	for i := range s.nums{
		j := i + rand.Intn(n - i)
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}
	return s.nums
}

func main() {
	nums := []int{1, 2, 3}
	solution := constructor(nums)
	res := solution.Shuffle()
	fmt.Println(res)
	res = solution.Reset()
	fmt.Println(res)
}
