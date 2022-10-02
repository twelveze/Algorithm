package main

import "fmt"

//直接遍历n会超时，n太大了。所以反其道而行之，假设一开始都是空排，每个空排最多可以坐两个家庭，那么就是2 * n,然后减去占用的排
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	var res int
	levelMap := make(map[int]int, 0)
	left := 0b11110000
	middle := 0b11000011
	right := 0b00001111
	for _, v := range reservedSeats {
		//因为1号座位和10号座位对这个排座没影响
		if v[1] != 1 && v[1] != 10 {
			//按位或
			levelMap[v[0]] |= 1 << (9 - v[1])
		}
	}
	//坐了人的排数
	seatWithPeople := len(levelMap)
	res = 2*n - seatWithPeople*2
	//看看坐了人的排可以坐几个家庭
	for _, v := range levelMap {
		//不等于0
		if (v&left) == 0 || (v&right) == 0 || (v&middle) == 0 {
			res++
		}
	}
	return res
}
func main() {
	reservedSeats := [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}
	res := maxNumberOfFamilies(3, reservedSeats)
	fmt.Println(res)
}
