package main

import (
	"Model"
	"fmt"
)
func DFSPreorder(root *Model.TreeNode) (res []*Model.TreeNode) {	//递归前序
	var preorder func(node *Model.TreeNode)
	preorder = func(node *Model.TreeNode) {
		if node == nil{
			return
		}
		res = append(res, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
func DFSInorder(root *Model.TreeNode) (res []*Model.TreeNode){	//递归中序
	var inorder func(node *Model.TreeNode)
	inorder = func(node *Model.TreeNode) {
		if node == nil{
			return
		}
		inorder(node.Left)
		res = append(res, node)
		inorder(node.Right)
	}
	inorder(root)
	return
}
func DFSPostorder(root *Model.TreeNode) (res []*Model.TreeNode) {	//递归后序
	var postorder func(node *Model.TreeNode)
	postorder = func(node *Model.TreeNode) {
		if node == nil{
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node)
	}
	postorder(root)
	return
}
func DFS(root *Model.TreeNode) (res []*Model.TreeNode){ //非递归的方式访问二叉树(前序遍历)
	var stack []*Model.TreeNode
	stack = append(stack, root)
	for len(stack) != 0{
		node := stack[len(stack) - 1]
		res = append(res, node)
		stack = stack[:len(stack) - 1]
		if node.Right != nil{
			stack = append(stack, node.Right)
		}
		if node.Left != nil{
			stack = append(stack, node.Left)
		}
	}
	return res
}
func DFSByNonRecursiveInorder(root *Model.TreeNode) (res []*Model.TreeNode){ //非递归的方式访问二叉树(中序遍历)
	var stack []*Model.TreeNode
	for root != nil || len(stack) > 0{
		for root != nil{
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, root)
		root = root.Right
	}
	return
}
func main() {
	n6 := &Model.TreeNode{Val: 11, Left: nil, Right: nil}
	n5 := &Model.TreeNode{Val: 7, Left: nil, Right: nil}
	n4 := &Model.TreeNode{Val: 15, Left: nil, Right: nil}
	n3 := &Model.TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &Model.TreeNode{Val: 9, Left: n6, Right: nil}
	n1 := &Model.TreeNode{Val: 3, Left: n2, Right: n3}
	//res := DFS(n1)
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
	for _, v := range roots{
		fmt.Printf("%d ",v.Val)
	}
}
