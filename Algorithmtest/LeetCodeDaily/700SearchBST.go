package main

import "Algorithm/Algorithmtest/Model"

func searchBST(root *Model.TreeNode, val int) *Model.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
func main() {

}
