package main

import (
	"fmt"
	"model"
)

//type TreeNode struct{
//	Val int
//	left *TreeNode
//	right *TreeNode
//}
func maxDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	n5 := &model.TreeNode{Val: 7, Left: nil, Right: nil}
	n4 := &model.TreeNode{Val: 15, Left: nil, Right: nil}
	n3 := &model.TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &model.TreeNode{Val: 9, Left: nil, Right: nil}
	n1 := &model.TreeNode{Val: 3, Left: n2, Right: n3}
	res := maxDepth(n1)
	fmt.Println(res)

}
