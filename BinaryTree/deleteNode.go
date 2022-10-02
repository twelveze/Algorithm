package main

import (
	"Model"
	"fmt"
)

func deleteNode(root *Model.TreeNode, target int) *Model.TreeNode{
	if root == nil{
		return nil
	}
	if target < root.Val{
		root.Left = deleteNode(root.Left, target)
		return root
	}
	if target > root.Val{
		root.Right = deleteNode(root.Right, target)
		return root
	}
	//到这一步就表明已经找到目标节点，准备删除
	//分为三种情况：1.节点无右子树，直接返回该节点的左孩子节点 2.节点无左子树，直接返回该节点的右孩子节点
	//3.节点同时存在左子树与右子树，找出子树中比当前节点小的最大节点替换自己(或比当前节点大的最小节点替换自己)
	if root.Right == nil{
		return root.Left
	}
	if root.Left == nil{
		return root.Right
	}
	minNode := root.Right
	for minNode.Left != nil{
		minNode = minNode.Left	//找到子树中比当前节点大的最小节点！
	}
	root.Val = minNode.Val
	root.Right = deleteMinNode(root.Right) 	//只需要改变右子树就好了，左子树不用变
	return root
}
func deleteMinNode(root *Model.TreeNode) *Model.TreeNode{
	if root.Left == nil{
		return root.Right
	}
	root.Left = deleteMinNode(root.Left)
	return root
}
func Visitfordelete(root *Model.TreeNode) (res []*Model.TreeNode){ //非递归的方式访问二叉树
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
func main() {
	n5 := &Model.TreeNode{Val: 7, Left: nil, Right: nil}
	n4 := &Model.TreeNode{Val: 15, Left: nil, Right: nil}
	n3 := &Model.TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &Model.TreeNode{Val: 9, Left: nil, Right: nil}
	n1 := &Model.TreeNode{Val: 3, Left: n2, Right: n3}
	root := deleteNode(n1, 15)
	res := Visitfordelete(root)
	for _, v := range res{
		fmt.Printf("%d",v.Val)
	}
}
