package main

import (
	"Model"
	"fmt"
)

func isBalanced(root *Model.TreeNode) bool{
	if root == nil{
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right){
		return false
	}
	leftHeight := Height(root.Left)
	rightHeight := Height(root.Right)
	if Model.Abs(leftHeight, rightHeight) > 1{
		return false
	}else{
		return true
	}
}
func Height(root *Model.TreeNode) int{
	if root == nil{
		return 0
	}
	return Model.Max(Height(root.Left), Height(root.Right)) + 1
}
func main() {
	n5 := &Model.TreeNode{Val: 7, Left: nil, Right: nil}
	n4 := &Model.TreeNode{Val: 15, Left: nil, Right: nil}
	n3 := &Model.TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &Model.TreeNode{Val: 9, Left: nil, Right: nil}
	n1 := &Model.TreeNode{Val: 3, Left: n2, Right: n3}
	res := isBalanced(n1)
	fmt.Println(res)
}