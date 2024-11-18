package main

import (
	"fmt"
	"model"
)

func reBuildBinaryTree(preorder []int, inorder []int) *model.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &model.TreeNode{Val: preorder[0], Left: nil, Right: nil}
	var index int
	for i, v := range inorder {
		if v == preorder[0] {
			index = i
		}
	}
	root.Left = reBuildBinaryTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = reBuildBinaryTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}
func Visit(root *model.TreeNode) (res []*model.TreeNode) { //非递归的方式访问二叉树
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
func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	res := reBuildBinaryTree(preorder, inorder)
	ans := Visit(res)
	for _, v := range ans {
		fmt.Print(" ", v.Val)
	}
}
