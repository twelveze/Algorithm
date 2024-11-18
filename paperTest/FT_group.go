//给定一个数组，不含重复元素，返回该数组所有可能的子集。如：
//输入：【1，2，3】
//输出：【
//【】，【1】，【2】，【3】，【1，2】，【1，3】，【2，3】，【1，2，3】
//】
package main

import "fmt"

func subset(arr []int) [][]int {
	lens := len(arr)
	var res [][]int
	visit := make(map[int]bool, 0)
	for i := 0; i < lens; i++ {
		var temp []int
		for j := 0; j < len(arr); j++ {
			visit[arr[j]] = false
		}
		res = dfsOfSubSet(visit, res, temp, i+1, arr, 0)
	}
	return res
}
func dfsOfSubSet(visit map[int]bool, res [][]int, temp []int, num int, arr []int, index int) [][]int {
	if num == len(temp) {
		tmp := make([]int, len(temp))
		copy(tmp, temp) //这里一定要值拷贝一下，不然后面改temp也会把res里的改掉
		res = append(res, tmp)
		return res
	}
	for i, v := range arr {
		if visit[v] == false && i >= index {
			temp = append(temp, v)
			visit[v] = true
			res = dfsOfSubSet(visit, res, temp, num, arr, i)
			temp = temp[:len(temp)-1]
			visit[v] = false
		}
	}
	return res
}
func main() {
	var arr = []int{1, 2, 3}
	res := subset(arr)
	fmt.Println(res)
}
