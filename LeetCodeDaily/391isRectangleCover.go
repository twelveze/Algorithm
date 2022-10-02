package main

import (
	"fmt"
	"math"
)

func isRectangleCover(rectangles [][]int) bool {
	// 计算每个小矩形面积是否等于大矩形面积
	// 看每个顶点出现的次数，如果最后出现一次的顶点不是四个，则说明不符合完美矩形
	area := 0	//所有小矩形相加的面积
	vertexMap := make(map[int]bool)
	//记录大矩形的左下角和右上角
	a1, b1 := math.MaxInt64, math.MaxInt64
	a2, b2 := math.MinInt64, math.MinInt64
	for _, rec := range rectangles{
		//小矩形的坐标
		x1 := rec[0]
		y1 := rec[1]
		x2 := rec[2]
		y2 := rec[3]
		//计算大矩形的左下角
		if x1 < a1 || y1 < b1{
			a1 = x1
			b1 = y1
		}
		//计算大矩形的右上角
		if x2 > a2 || y2 > b2{
			a2 = x2
			b2 = y2
		}
		area += (x2 - x1) * (y2 - y1)
		record(vertexMap, x1, y1)
		record(vertexMap, x1, y2)
		record(vertexMap, x2, y1)
		record(vertexMap, x2, y2)
	}
	totalArea := (a2 - a1) * (b2 - b1)
	if totalArea != area{
		return false
	}
	//count记录出现奇数次的顶点个数
	var count int
	for i := range vertexMap{
		if vertexMap[i] == true{
			count++
		}
	}
	//四个出现奇数次的顶点正好是大矩形的四个顶点
	return count == 4 && vertexMap[getKey(a1, b1)] && vertexMap[getKey(a1, b2)] &&
		vertexMap[getKey(a2, b1)] && vertexMap[getKey(a2, b2)]
}
func record(vertexMap map[int]bool, x, y int){
	//记录顶点出现的次数，如果一个顶点出现偶数次，则记录为false
	var key int
	key = getKey(x, y)
	if vertexMap[key]{
		vertexMap[key] = false
	}else{
		vertexMap[key] = true
	}
}
func getKey(x, y int) int{
	// 二维坐标转一维，方便比较
	// 100000007是随便取的一个大质数
	// 这里即使溢出了也没什么问题
	return x * 100000007 + y
}
func main() {
	var rectangles = [][]int{{1,1,3,3},{3,1,4,2},{3,2,4,4},{1,3,2,4},{2,3,3,4}}
	res := isRectangleCover(rectangles)
	fmt.Println(res)
}
