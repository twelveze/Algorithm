package main

import (
	"fmt"
	"sort"
)

func Permutation(seed []byte, size int) []string{
	visited := make(map[byte]bool)
	sort.Slice(seed, func(i, j int) bool {
		return seed[i] < seed[j]
	})
	var res []string
	var temp []byte
	for i := 0; i < len(seed); i++{
		visited[seed[i]] = false
	}
	res = backTrack(seed, temp, size, visited, res)
	return res
}
func backTrack(seed []byte, temp []byte, size int, visited map[byte]bool, res []string) []string{
	if len(temp) == size{
		var str string
		str = string(temp[:])
		res = append(res, str)
		return res
	}
	for _, b := range seed{
		if visited[b] == false{
			temp = append(temp, b)
			visited[b] = true
			res = backTrack(seed, temp, size, visited, res)
			temp = temp[:len(temp) - 1]
			visited[b] = false
		}
	}
	return res
}
func main() {
	var seed = []byte{'0','1','2'}
	size := 2
	res := Permutation(seed,size)
	fmt.Println(res)
}
