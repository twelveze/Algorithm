package main

import (
	"fmt"
	"model"
)

// leetcode - 687 最长同值路径
var ans = 0

func dfs(root *model.TreeNode) int {
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
func longestUnivaluePath(root *model.TreeNode) int {
	dfs(root)
	return ans
}
func main() {
	node6 := &model.TreeNode{Val: 5, Left: nil, Right: nil}
	node5 := &model.TreeNode{Val: 4, Left: nil, Right: nil}
	node4 := &model.TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &model.TreeNode{Val: 5, Left: nil, Right: node6}
	node2 := &model.TreeNode{Val: 4, Left: node4, Right: node5}
	node1 := &model.TreeNode{Val: 1, Left: node2, Right: node3}

	res := longestUnivaluePath(node1)
	fmt.Println(res)
}
