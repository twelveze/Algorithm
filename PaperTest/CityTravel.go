package main

import "fmt"

func CityTravel(start, end int, open, rest []int, roads [][]int,)int{
	totaltime := 0
	firstcity := start
	for firstcity < end{
		for i := 0; i < len(roads); i++{
			if roads[i][0] == firstcity {
				totaltime += roads[i][2]
				totaltime += (totaltime / open[roads[i][1]] + 1) * open[roads[i][1]]
				totaltime += rest[roads[i][1]]
				break
			}
			firstcity = roads[i][1]
		}
	}
	return totaltime
}
func main() {
	var n, m int
	open := make([]int, 0)
	rest := make([]int, 0)
	fmt.Scanf("%d %d",&n, &m)
	roads := make([][]int, m)
	for i := 0; i < n; i++{
		var w int
		fmt.Scan(&w)
		open = append(open, w)
	}
	for i := 0; i < n; i++{
		var p int
		fmt.Scan(&p)
		rest = append(rest, p)
	}
	for i := 0; i < m; i++{
		road := make([]int, 3)
		var start int
		var end int
		var len int
		fmt.Scanf("%d %d %d",&start, &end, &len)
		road = append(road, start, end, len)
		roads = append(roads, road)
	}
	var start int
	var end int
	fmt.Scanf("%d %d", &start, &end)
	res := CityTravel(start, end, open, rest, roads)
	fmt.Println(res)
}
