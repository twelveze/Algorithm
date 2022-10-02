package main

import (
	"Model"
	"fmt"
)

func searchBST(root *Model.TreeNode, target int) *Model.TreeNode { // 递归
	if root == nil {
		return nil
	}
	if root.Val > target {
		return searchBST(root.Left, target)
	} else if root.Val < target {
		return searchBST(root.Right, target)
	} else {
		return root
	}
}
func searchBSTdd(root *Model.TreeNode, target int) *Model.TreeNode { //迭代
	for root != nil {
		if root.Val == target {
			return root
		} else if root.Val > target {
			root = root.Left
		} else if root.Val < target {
			root = root.Right
		}
	}
	return nil
}

func main() {
	n5 := &Model.TreeNode{Val: 7, Left: nil, Right: nil}
	n4 := &Model.TreeNode{Val: 15, Left: nil, Right: nil}
	n3 := &Model.TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &Model.TreeNode{Val: 9, Left: nil, Right: nil}
	n1 := &Model.TreeNode{Val: 3, Left: n2, Right: n3}
	res := searchBST(n1, 15)
	fmt.Println(res.Val)
}
