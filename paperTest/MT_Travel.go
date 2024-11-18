package main

import "fmt"

func MT_Travel() {

}
func main() {
	var T int
	records := make([][]int, T)
	for i := 0; i < T; i++ {
		var n, m, k int
		record := make([]int, 0)
		fmt.Scanf("%d %d %d", &n, &m, &k)
		record = append(record, n)
		record = append(record, m)
		record = append(record, k)
		records[i] = record
		for j := 0; j < m; j++ {
			var u, v int
			fmt.Scanf("%d %d", &u, &v)

		}
	}

}
