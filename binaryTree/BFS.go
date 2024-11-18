package main

import (
	"container/list"
	"fmt"
	"model"
)

func BFS(root *model.TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if level == len(res) {
		res = append(res, []int{root.Val}) //还未初始化第level层的数组，初始化并赋值
	} else {
		res[level] = append(res[level], root.Val) //这里表示初始化第level层的数组了，直接append该层数组
	}
	res = BFS(root.Left, level+1, res)
	res = BFS(root.Right, level+1, res)
	return res
}
func BFSbyQueue(root *model.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		var cur []int
		listLength := queue.Len()
		for i := 0; i < listLength; i++ {
			node := queue.Remove(queue.Front()).(*model.TreeNode) //层次遍历，从左向右依次取
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
func BFSbyFakeQueue(root *model.TreeNode) [][]int {
	var res [][]int
	tempQueue := make([]*model.TreeNode, 0)
	tempQueue = append(tempQueue, root)
	for len(tempQueue) > 0 {
		var cur []int
		queueLen := len(tempQueue)
		for i := 0; i < queueLen; i++ {
			cur = append(cur, tempQueue[0].Val)
			if tempQueue[0].Left != nil {
				tempQueue = append(tempQueue, tempQueue[0].Left)
			}
			if tempQueue[0].Right != nil {
				tempQueue = append(tempQueue, tempQueue[0].Right)
			}
			tempQueue = tempQueue[1:]
		}
		res = append(res, cur)
	}
	return res
}
func main() {
	n5 := &model.TreeNode{Val: 7, Left: nil, Right: nil}
	n4 := &model.TreeNode{Val: 15, Left: nil, Right: nil}
	n3 := &model.TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &model.TreeNode{Val: 9, Left: nil, Right: nil}
	n1 := &model.TreeNode{Val: 3, Left: n2, Right: n3}
	res := BFS(n1, 0, [][]int{})
	resByQueue := BFSbyQueue(n1)
	resByFakeQueue := BFSbyFakeQueue(n1)
	fmt.Println(res)
	fmt.Println(resByQueue)
	fmt.Println(resByFakeQueue)
}
