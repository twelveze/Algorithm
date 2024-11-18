package main

import (
	"fmt"
	"math/rand"
)

type SolutionFor519 struct {
	m, n, total int
	mp          map[int]int
}

func ConstructorFor519(m int, n int) SolutionFor519 {
	return SolutionFor519{m: m, n: n, total: m * n, mp: map[int]int{}}
}

func (s *SolutionFor519) Flip() (ans []int) {
	rand := rand.Intn(s.total)
	s.total--
	if y, ok := s.mp[rand]; ok {
		ans = []int{y / s.n, y % s.n}
	} else {
		ans = []int{rand / s.n, rand % s.n}
	}
	if y, ok := s.mp[s.total]; ok {
		s.mp[rand] = y
	} else {
		s.mp[rand] = s.total
	}
	return
}

func (s *SolutionFor519) Reset() {
	s.total = s.m * s.n
	s.mp = map[int]int{}
}

func main() {
	sol := ConstructorFor519(3, 1)
	res := sol.Flip()
	fmt.Println(res)
	res = sol.Flip()
	fmt.Println(res)
	res = sol.Flip()
	fmt.Println(res)
	sol.Reset()
}
