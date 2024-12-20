package main

import (
	"fmt"
	"model"
)

//func maxDepth(root *model.Node) int{
//	if root == nil{
//		return 0
//	}
//	max := 0
//	for _, n := range root.Children{
//		temp := maxDepth(n)
//		if temp > max{
//			max = temp
//		}
//	}
//	return max + 1
//}

func maxDepth(root *model.Node) int {
	if root == nil {
		return 0
	}
	return getDepth(root) + 1
}

func getDepth(root *model.Node) (res int) {
	max := 0
	if root == nil {
		return max
	} else {
		max = 1
	}
	for _, child := range root.Children {
		temp := max + getDepth(child)
		if temp > res {
			res = temp
		}
	}
	return
}
func main() {
	n6 := &model.Node{Val: 6, Children: nil}
	n5 := &model.Node{Val: 5, Children: nil}
	n4 := &model.Node{Val: 4, Children: nil}
	n3 := &model.Node{Val: 2, Children: nil}
	n2 := &model.Node{Val: 3, Children: []*model.Node{n5, n6}}
	n1 := &model.Node{Val: 1, Children: []*model.Node{n2, n3, n4}}
	res := maxDepth(n1)
	fmt.Println(res)
}
