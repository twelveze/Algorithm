package main

import "fmt"

func WholeSort(nums []int) [][]int {
	var temp []int
	var ans [][]int
	visited := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		visited[nums[i]] = false
	}
	ans = dfs(nums, ans, temp, visited)
	return ans
}
func dfs(nums []int, ans [][]int, temp []int, visited map[int]bool) [][]int {
	if len(temp) == len(nums) {
		//为什么加入解集时，要将数组内容拷贝到一个新的数组里，再加入解集？

		//因为该temp变量存的是地址引用，结束当前递归时，将它加入 ans 后，该算法还要进入别的递归分支继续搜索，
		//还要继续将这个 path 传给别的递归调用，它所指向的内存空间还要继续被操作，所以 ans 中的 temp 的内容会被改变，这就不对。
		tempCopy := make([]int, len(temp))
		copy(tempCopy, temp)
		ans = append(ans, tempCopy)
		return ans
	}
	for _, num := range nums {
		if visited[num] == false {
			temp = append(temp, num)
			visited[num] = true
			ans = dfs(nums, ans, temp, visited)
			temp = temp[:len(temp)-1]
			visited[num] = false
		}
	}
	return ans
}
func main() {
	nums := []int{5, 4, 6, 2}
	res := WholeSort(nums)
	fmt.Println(res)
}
