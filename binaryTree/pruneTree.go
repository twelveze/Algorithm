package main

import "model"

func pruneTree(root *model.TreeNode) *model.TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Right == nil && root.Left == nil && root.Val == 0 {
		return nil
	}
	return root
}
