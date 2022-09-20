package main

import "fmt"

func findContinuousSequence(target int) [][]int{
	if target < 3{
		return [][]int{}
	}
	arr := make([]int, target)
	for i := 1; i < target; i++{
		arr[i] = i
	}
	res := make([][]int, 0)
	start, end := 1,2
	sum := arr[start] + arr[end]
	for end <= target / 2 + 1{
		if sum < target{
			end++
			sum += arr[end]
		}else if sum > target{
			sum -= arr[start]
			start++
		}else{
			temp := arr[start:end + 1]
			res = append(res, temp)
			end++
			sum += arr[end]
		}
	}
	return res
}
func main() {
	target := 9
	res := findContinuousSequence(target)
	fmt.Println(res)
}
