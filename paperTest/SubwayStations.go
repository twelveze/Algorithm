package main

import "fmt"

func SubwayStation(cross, change [][]int, n int) []int {
	for i := 0; i < len(change); i++ {
		for j := 0; j < len(cross); j++ {
			cross[j] = changeStation(cross[j], change[i])
		}
	}
	res := make([]int, n)
	for _, v := range cross {
		for _, val := range v {
			res[val-1]++
		}
	}
	return res
}
func changeStation(cross []int, change []int) []int {
	for i := 0; i <= 1; i++ {
		if cross[i] == change[0] {
			cross[i] = change[1]
		} else if cross[i] == change[1] {
			cross[i] = change[0]
		}
	}
	return cross
}
func main() {
	var n, m, q int
	_, err := fmt.Scanf("%d %d %d", &n, &m, &q)
	if err != nil {
	} else {
		cross := make([][]int, m)
		change := make([][]int, q)
		for i := 0; i < m; i++ {
			var temp1, temp2 int
			fmt.Scanln(&temp1, &temp2)
			cross[i] = append(cross[i], temp1)
			cross[i] = append(cross[i], temp2)
		}
		for i := 0; i < q; i++ {
			var temp1, temp2 int
			fmt.Scanln(&temp1, &temp2)
			change[i] = append(change[i], temp1)
			change[i] = append(change[i], temp2)
		}
		res := SubwayStation(cross, change, n)
		for _, v := range res {
			fmt.Printf("%d ", v)
		}
	}
}
