package main

import (
	"fmt"
	"model"
)

func levelOrder(root *model.Node) [][]int {
	var res [][]int
	var node []*model.Node
	if root == nil {
		return [][]int{}
	} else {
		node = append(node, root)
		res = append(res, []int{root.Val})
	}
	for len(node) != 0 {
		var tempNode []*model.Node
		var tempRes []int
		for i := 0; i < len(node); i++ { //目前这一层的所有节点
			for j := 0; j < len(node[i].Children); j++ { //存储临时结果，然后把节点的子节点一个个加入到下一层的遍历队列
				tempNode = append(tempNode, node[i].Children[j])
				tempRes = append(tempRes, node[i].Children[j].Val)
			}
		}
		node = tempNode
		res = append(res, tempRes)
	}
	n := len(res)
	return res[:n-1]
}

func levelOrderByAnswer(root *model.Node) (ans [][]int) {
	if root == nil {
		return [][]int{}
	}
	queue := []*model.Node{root}
	for queue != nil {
		var level []int
		temp := queue
		queue = nil
		for _, node := range temp {
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		ans = append(ans, level)
	}
	return
}

func main() {
	n6 := model.Node{Val: 6, Children: nil}
	n5 := model.Node{Val: 5, Children: nil}
	n4 := model.Node{Val: 4, Children: nil}
	n3 := model.Node{Val: 3, Children: []*model.Node{&n5, &n6}}
	n2 := model.Node{Val: 2, Children: nil}
	n1 := model.Node{Val: 1, Children: []*model.Node{&n3, &n2, &n4}}
	res := levelOrder(&n1)
	fmt.Println(res)
}
