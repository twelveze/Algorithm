package main

import (
	"fmt"
	"model"
)

//O(1)复杂度删除一个节点
func deleteNode(node *model.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	n3 := &model.ListNode{Val: 3, Next: nil}
	n2 := &model.ListNode{Val: 2, Next: n3}
	n1 := &model.ListNode{Val: 1, Next: n2}
	deleteNode(n2)
	for n1 != nil {
		fmt.Println(n1.Val)
		n1 = n1.Next
	}
}
