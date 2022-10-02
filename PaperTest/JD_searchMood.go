package main

import "fmt"

func searchMood(mood []int) int{ //Accept
	max := 0
	min := 10001
	moodMap := make(map[int]int,0)
	for i := 0; i < len(mood); i++{
		if mood[i] <= min{
			min = mood[i]
			moodMap[0] = i + 1
			if len(mood) - i < moodMap[0]{
				moodMap[0] = i + 1
			}
		}
		if mood[i] >= max{
			max = mood[i]
			moodMap[1] = i + 1
			if len(mood) - i < moodMap[1]{
				moodMap[1] = i + 1
			}
		}
	}
	var large int
	var small int
	if moodMap[0] > moodMap[1]{
		large = moodMap[0]
		small = moodMap[1]
	}else{
		large = moodMap[1]
		small = moodMap[0]
	}
	gap := minn(large, len(mood) - small + 1)
	if gap > small + (len(mood) - large + 1){
		return small + (len(mood) - large + 1)
	}else{
		return gap
	}
}
func minn(a, b int) int{
	if a > b{
		return b
	}else{
		return a
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	mood := make([]int, 0)
	for i := 0; i < n; i++{
		var temp int
		fmt.Scan(&temp)
		mood = append(mood, temp)
	}
	res := searchMood(mood)
	fmt.Println(res)
}
