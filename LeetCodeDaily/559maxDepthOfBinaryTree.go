package main

import (
	"Algorithm/Algorithmtest/Model"
	"fmt"
)

//func maxDepth(root *Model.Node) int{
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

func maxDepth(root *Model.Node) int {
	if root == nil {
		return 0
	}
	return getDepth(root) + 1
}

func getDepth(root *Model.Node) (res int) {
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
	n6 := &Model.Node{Val: 6, Children: nil}
	n5 := &Model.Node{Val: 5, Children: nil}
	n4 := &Model.Node{Val: 4, Children: nil}
	n3 := &Model.Node{Val: 2, Children: nil}
	n2 := &Model.Node{Val: 3, Children: []*Model.Node{n5, n6}}
	n1 := &Model.Node{Val: 1, Children: []*Model.Node{n2, n3, n4}}
	res := maxDepth(n1)
	fmt.Println(res)
}
