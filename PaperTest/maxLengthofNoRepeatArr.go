package main

import "fmt"

func maxLengthOfNoRepeatArr(arr []int)int{
	arrmap := map[int]int{}
	arrLen := len(arr)
	temp := 0
	res := 0
	left := 0
	for i := 0; i < arrLen; i++{
		if id, ok := arrmap[arr[i]]; !ok{
			temp++
		}else{
			if id > left{
				left = id
			}
			temp = i - left
		}
		arrmap[arr[i]] = i
		if res < temp{
			res = temp
		}
	}
	return res
}
func maxLengthOfNoRepeatArr2(arr []int) int{
	arrmap := map[int]int{}
	arrLen := len(arr)
	temp := 0
	index := 0
	res := 0
	for i := 0; i < arrLen; i++{
		if _, ok := arrmap[arr[i]]; ok{
			temp = i - index
			index = arrmap[arr[i]] + 1
			arrmap[arr[i]] = i
		}else{
			arrmap[arr[i]] = i
			continue
		}
		if temp > res{
			res = temp
		}
	}
	return res
}
func main() {
	arr := []int{2,2,3,4,5,3}
	res := maxLengthOfNoRepeatArr(arr)
	res1 := maxLengthOfNoRepeatArr2(arr)
	fmt.Println(res)
	fmt.Println(res1)
}
