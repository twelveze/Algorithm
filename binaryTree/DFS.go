package main

import (
	"fmt"
	"model"
)

func DFSPreorder(root *model.TreeNode) (res []*model.TreeNode) { //递归前序
	var preorder func(node *model.TreeNode)
	preorder = func(node *model.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
func DFSInorder(root *model.TreeNode) (res []*model.TreeNode) { //递归中序
	var inorder func(node *model.TreeNode)
	inorder = func(node *model.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node)
		inorder(node.Right)
	}
	inorder(root)
	return
}
func DFSPostorder(root *model.TreeNode) (res []*model.TreeNode) { //递归后序
	var postorder func(node *model.TreeNode)
	postorder = func(node *model.TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node)
	}
	postorder(root)
	return
}
func DFS(root *model.TreeNode) (res []*model.TreeNode) { //非递归的方式访问二叉树(前序遍历)
	var stack []*model.TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		res = append(res, node)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
func DFSByNonRecursiveInorder(root *model.TreeNode) (res []*model.TreeNode) { //非递归的方式访问二叉树(中序遍历)
	var stack []*model.TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root)
		root = root.Right
	}
	return
}
func main() {
	n6 := &model.TreeNode{Val: 11, Left: nil, Right: nil}
	n5 := &model.TreeNode{Val: 7, Left: nil, Right: nil}
	n4 := &model.TreeNode{Val: 15, Left: nil, Right: nil}
	n3 := &model.TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &model.TreeNode{Val: 9, Left: n6, Right: nil}
	n1 := &model.TreeNode{Val: 3, Left: n2, Right: n3}
	//res := dfs(n1)
	//for _, v := range res{
	//	fmt.Printf("%d ",v.Val)
	//}
	//fmt.Println()
	//roots := DFSPreorder(n1)
	//for _, v := range roots{
	//	fmt.Printf("%d ",v.Val)
	//}
	//fmt.Println()
	//roots = DFSInorder(n1)
	//for _, v := range roots{
	//	fmt.Printf("%d ",v.Val)
	//}
	//fmt.Println()
	//roots = DFSPostorder(n1)
	//for _, v := range roots{
	//	fmt.Printf("%d ",v.Val)
	//}
	//fmt.Println()
	roots := DFSByNonRecursiveInorder(n1)
	for _, v := range roots {
		fmt.Printf("%d ", v.Val)
	}
}
