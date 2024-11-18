package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	startTime := timeSeries[0]
	endTime := timeSeries[0]
	var res int
	for i := 0; i < len(timeSeries); i++ {
		if timeSeries[i] > endTime {
			temp := endTime - startTime + 1
			res = res + temp
			startTime = timeSeries[i]
			endTime = timeSeries[i] + duration - 1
		} else {
			endTime = timeSeries[i] + duration - 1
		}
	}
	res = res + endTime - startTime + 1
	return res
}
func main() {
	var timeSeries = []int{1, 2}
	duration := 2
	res := findPoisonedDuration(timeSeries, duration)
	fmt.Println(res)
}
