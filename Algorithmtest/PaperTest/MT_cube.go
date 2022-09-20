package main

import (
	"fmt"
	"sort"
)

func cube(T int, cubes, locations [][]int) []string{
	res := make([]string, T)
	for i := 0; i < T; i++{
		sum := 0
		sort.Ints(locations[i])
		for j := len(locations[i]) - 1; j >= 0; j--{
			if cubes[i][1] > 0{
				sum += locations[i][j] * locations[i][j] * locations[i][j]
				cubes[i][1]--
			}else{
				sum += locations[i][j]
			}
		}
		if sum < cubes[i][2]{
			res[i] = "NO"
		}else{
			res[i] = "YES"
		}
	}
	return res
}
func main() {
	var T int
	fmt.Scan(&T)
	var cubes = make([][]int, T)
	var locations = make([][]int, T)
	for i:= 0; i < T; i++{
		var n, m, val int
		var cube []int
		fmt.Scanf("%d %d %d", &n,&m, &val)
		cube = append(cube, n, m, val)
		cubes[i] = cube
		locates := make([]int, 0)
		for j := 0; j < n; j++{
			var locate int
			fmt.Scan(&locate)
			locates = append(locates, locate)
		}
		locations[i] = locates
	}
	res := cube(T, cubes, locations)
	for i := 0; i < T; i++{
		fmt.Println(res[i])
	}
}
