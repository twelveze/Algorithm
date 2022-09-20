package main

import (
	"Model"
	"fmt"
)

//func findTilt(root *Model.TreeNode) (ans int) { //深度优先遍历
//	var dfs func(*Model.TreeNode) int
//	dfs = func(node *Model.TreeNode) int{
//		if node == nil{
//			return 0
//		}
//		LeftSum := dfs(node.Left)
//		RightSum := dfs(node.Right)
//		ans += Model.Abs(LeftSum, RightSum)
//		return LeftSum + RightSum + node.Val
//	}
//	dfs(root)
//	return
//}
func findTilt(root *Model.TreeNode) int{ //也是递归
	if root == nil{
		return 0
	}
	res := findTilt(root.Left) + findTilt(root.Right) + Model.Abs(getSum(root.Left), getSum(root.Right))
	return res
}
func getSum(root *Model.TreeNode) int{
	if root == nil{
		return 0
	}
	return getSum(root.Left) + getSum(root.Right) + root.Val
}
func main() {
	n6 := &Model.TreeNode{Val: 7, Left: nil, Right: nil}
	n5 := &Model.TreeNode{Val: 5, Left: nil, Right: nil}
	n4 := &Model.TreeNode{Val: 3, Left: nil, Right: nil}
	n3 := &Model.TreeNode{Val: 9, Left: nil, Right: n6}
	n2 := &Model.TreeNode{Val: 2, Left: n4, Right: n5}
	n1 := &Model.TreeNode{Val: 4, Left: n2, Right: n3}
	res := findTilt(n1)
	fmt.Println(res)
}
