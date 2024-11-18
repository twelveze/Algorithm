package main

import (
	"fmt"
	"model"
)

// leetcode - 19 删除链表的倒数第N个节点

// 双指针,头指针先走n步，走到null的时候，慢指针的下一个就是要删除的节点
func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
	dummy := &model.ListNode{Next: head}
	first, second := head, dummy
	// 先走n步
	for i := 0; i < n; i++ {
		first = first.Next
	}

	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	n5 := &model.ListNode{Val: 5, Next: nil}
	n4 := &model.ListNode{Val: 4, Next: n5}
	n3 := &model.ListNode{Val: 3, Next: n4}
	n2 := &model.ListNode{Val: 2, Next: n3}
	head := &model.ListNode{Val: 1, Next: n2}
	head = removeNthFromEnd(head, 2)
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
