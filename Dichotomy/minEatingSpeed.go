package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, hour int) int{
	maxSpeed := 1
	for i := 0; i < len(piles); i++{
		if piles[i] > maxSpeed{
			maxSpeed = piles[i]
		}
	}
	left := 1
	right := maxSpeed
	for left < right{
		mid := (right + left) / 2
		if ableEat(piles, mid, hour){
			left = mid + 1
		}else{
			right = mid
		}
	}
	return left
}
func ableEat(piles []int, speed, hour int) bool{
	var sumhour float64
	for i := 0; i < len(piles); i++{
		sumhour = sumhour + math.Ceil(float64(piles[i]) / float64(speed))
	}
	return sumhour > float64(hour)
}
func main() {
	var n int
	fmt.Scanf("%d",&n)
	var piles []int
	for i := 0; i < n; i++  {
		var temp int
		fmt.Scanf("%d",&temp)
		piles = append(piles, temp)
	}
	var hour int
	fmt.Scanf("%d",&hour)
	res := minEatingSpeed(piles, hour)
	fmt.Println(res)
}
