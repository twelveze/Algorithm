package main

import "fmt"

func xipaiI(arr []int) []int{
	lens := len(arr)
	front := make([]int, 0)
	behind := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < lens / 2; i++{
		front = append(front, arr[i])
	}
	for i := lens / 2; i < lens; i++{
		behind = append(behind, arr[i])
	}
	flag := 1
	frontIndex := 0
	behindIndex := 0
	for lens > 0{
		if flag == 1{
			flag = -1
			res = append(res, behind[behindIndex])
			behindIndex++
			lens--
		}else{
			flag = 1
			res = append(res, front[frontIndex])
			frontIndex++
			lens--
		}
	}
	return res
}
func main() {
	var arr = []int{1,2,3,4,5,6}
	res := xipaiI(arr)
	fmt.Println(res)
}
