package main

import (
	"fmt"
)

func twoSum(arr []int, target int)[]int{
	res := make([]int, 0)
	arrMap := make(map[int]int,0)
	//for i := 0; i < len(arr); i++{
	//	arrMap[arr[i]] = i
	//}
	//for i := 0; i < len(arr); i++{
	//	if arr[i] * 2 == target{
	//		continue
	//	}
	//	if v, ok := arrMap[target - arr[i]]; ok{
	//		res = append(res, i, v)
	//		return res
	//	}
	//}
	arrMap[arr[0]] = 0
	for i := 1; i < len(arr); i++{
		if v, ok := arrMap[target - arr[i]]; ok{
			res = append(res, i, v)
			return res
		}else{
			arrMap[arr[i]] = i
		}
	}
	return res
}
func main() {
	var arr = []int{2,3,1,4}
	target := 4
	res := twoSum(arr, target)
	fmt.Println(res)
}
