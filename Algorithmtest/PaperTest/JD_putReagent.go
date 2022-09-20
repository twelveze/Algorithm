package main

import "fmt"

func react(reaction [][]int, n int) int {
	record := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		record[i] = make([]int, n+1)
	}
	for i := 0; i < len(reaction); i++ {
		record[reaction[i][0]][reaction[i][1]] = reaction[i][2]
		record[reaction[i][1]][reaction[i][0]] = reaction[i][2]
	}
	return len(record)
}
func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	reaction := make([][]int, 0)
	for i := 0; i < m; i++ {
		var a, b, c int
		temp := make([]int, 0)
		fmt.Scanf("%d %d %d", &a, &b, &c)
		temp = append(temp, a, b, c)
		reaction = append(reaction, temp)
	}
	res := react(reaction, n)
	fmt.Println(res)
}
