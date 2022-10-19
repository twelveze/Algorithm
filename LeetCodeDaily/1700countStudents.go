package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	count0 := 0
	count1 := 0
	//分别计算出学生中喜欢0和喜欢1的个数
	for _, v := range students {
		if v == 1 {
			count1++
		} else {
			count0++
		}
	}
	for _, v := range sandwiches {
		if v == 0 && count0 > 0 {
			count0--
		} else if v == 1 && count1 > 0 {
			count1--
		} else {
			break
		}
	}
	return count1 + count0
}

func main() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	res := countStudents(students, sandwiches)
	fmt.Println(res)
}
