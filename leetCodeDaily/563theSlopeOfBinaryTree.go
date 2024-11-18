package main

import (
	"fmt"
	"model"
)

//func findTilt(root *model.TreeNode) (ans int) { //深度优先遍历
//	var dfs func(*model.TreeNode) int
//	dfs = func(node *model.TreeNode) int{
//		if node == nil{
//			return 0
//		}
//		LeftSum := dfs(node.Left)
//		RightSum := dfs(node.Right)
//		ans += model.Abs(LeftSum, RightSum)
//		return LeftSum + RightSum + node.Val
//	}
//	dfs(root)
//	return
//}
func findTilt(root *model.TreeNode) int { //也是递归
	if root == nil {
		return 0
	}
	res := findTilt(root.Left) + findTilt(root.Right) + model.Abs(getSum(root.Left), getSum(root.Right))
	return res
}
func getSum(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	return getSum(root.Left) + getSum(root.Right) + root.Val
}
func main() {
	n6 := &model.TreeNode{Val: 7, Left: nil, Right: nil}
	n5 := &model.TreeNode{Val: 5, Left: nil, Right: nil}
	n4 := &model.TreeNode{Val: 3, Left: nil, Right: nil}
	n3 := &model.TreeNode{Val: 9, Left: nil, Right: n6}
	n2 := &model.TreeNode{Val: 2, Left: n4, Right: n5}
	n1 := &model.TreeNode{Val: 4, Left: n2, Right: n3}
	res := findTilt(n1)
	fmt.Println(res)
}
