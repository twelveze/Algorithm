package main

import (
	"Algorithm/Algorithmtest/Model"
	"fmt"
)

//O(1)复杂度删除一个节点
func deleteNode(node *Model.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	n3 := &Model.ListNode{Val: 3, Next: nil}
	n2 := &Model.ListNode{Val: 2, Next: n3}
	n1 := &Model.ListNode{Val: 1, Next: n2}
	deleteNode(n2)
	for n1 != nil {
		fmt.Println(n1.Val)
		n1 = n1.Next
	}
}
