package main

import (
	"fmt"
)

func main() {
	//arr := []int{1, 2, 1, 3, 4, 3, 5, 5, 4}
	//res := 0
	//for _, v := range arr {
	//	res ^= v
	//}
	//fmt.Println(res)

	//变种:一个数组内数字两两相同,其中只有两个数字是不一样的,找出这两个数字
	//https://www.cnblogs.com/lilideng/p/FindDifference.html
	arr := []int{1, 2, 3, 1, 2, 3, -7, 6}
	res := 0
	for _, v := range arr {
		res ^= v
	}
	count := 0 //二进制中第几位是1
	for !((res & 1) == 1) {
		res = res >> 1
		count++
	}
	var temp1 []int
	var temp2 []int
	for _, v := range arr {
		if (v >> count & 1) == 1 {
			temp1 = append(temp1, v)
		} else {
			temp2 = append(temp2, v)
		}
	}
	res1 := 0
	res2 := 0
	for _, v := range temp1 {
		res1 ^= v
	}
	for _, v := range temp2 {
		res2 ^= v
	}
	fmt.Println(res1, res2)
	//fmt.Println(6>>1, 7>>1)
	fmt.Println(count)
}
