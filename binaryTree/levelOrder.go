package main

import (
	"container/list"
	"fmt"
	"model"
)

func levelOrder(node *model.TreeNode) [][]int {
	var res [][]int
	var levelRes []*model.TreeNode
	levelRes = append(levelRes, node)
	for len(levelRes) > 0 {
		var cur []int
		length := len(levelRes)
		for i := 0; i < length; i++ {
			if levelRes[0].Left != nil {
				levelRes = append(levelRes, levelRes[0].Left)
			}
			if levelRes[0].Right != nil {
				levelRes = append(levelRes, levelRes[0].Right)
			}
			cur = append(cur, levelRes[0].Val)
			levelRes = levelRes[1:]
		}
		res = append(res, cur)
	}
	return res
}
func levelOrderByList(node *model.TreeNode) [][]int {
	var res [][]int
	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		var cur []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*model.TreeNode)
			cur = append(cur, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, cur)
	}
	return res
}

//层序遍历的实践！
func main() {
	n5 := &model.TreeNode{Val: 7, Left: nil, Right: nil}
	n4 := &model.TreeNode{Val: 15, Left: nil, Right: nil}
	n3 := &model.TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &model.TreeNode{Val: 9, Left: nil, Right: nil}
	n1 := &model.TreeNode{Val: 3, Left: n2, Right: n3}
	res := levelOrder(n1)
	res1 := levelOrderByList(n1)
	fmt.Println(res)
	fmt.Println(res1)
}
