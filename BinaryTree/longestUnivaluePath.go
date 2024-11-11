package main

import (
	"Model"
	"fmt"
)

// leetcode - 687 最长同值路径
var ans = 0

func dfs(root *Model.TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)

	leftTmp, rightTmp := 0, 0
	if root.Left != nil && root.Left.Val == root.Val {
		leftTmp = left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		rightTmp = right + 1
	}
	if leftTmp+rightTmp > ans {
		ans = leftTmp + rightTmp
	}

	if leftTmp > rightTmp {
		return leftTmp
	} else {
		return rightTmp
	}
}
func longestUnivaluePath(root *Model.TreeNode) int {
	dfs(root)
	return ans
}
func main() {
	node6 := &Model.TreeNode{Val: 5, Left: nil, Right: nil}
	node5 := &Model.TreeNode{Val: 4, Left: nil, Right: nil}
	node4 := &Model.TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &Model.TreeNode{Val: 5, Left: nil, Right: node6}
	node2 := &Model.TreeNode{Val: 4, Left: node4, Right: node5}
	node1 := &Model.TreeNode{Val: 1, Left: node2, Right: node3}

	res := longestUnivaluePath(node1)
	fmt.Println(res)
}
