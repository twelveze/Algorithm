package main

import (
	"fmt"
)

//type listNode struct {
//	Val int
//	Next *listNode
//}
func cycleByHash(head *listNode) bool { //利于hash表解决
	m := make(map[*listNode]bool)
	for head != nil {
		if _, exist := m[head]; exist {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}
func cycleByFastLowPointer(head *listNode) bool { //使用快慢指针的方法
	if head == nil {
		return false
	}
	fast := head.Next //快指针，一次走两个节点
	for fast != nil && head != nil {
		if fast == head { //快慢指针相遇，表示有环
			return true
		}
		fast = fast.Next.Next
		head = head.Next
	}
	return false
}
func main() {
	n4 := &listNode{Val: -4, Next: nil}
	n3 := &listNode{Val: 0, Next: n4}
	n2 := &listNode{Val: 2, Next: n3}
	n4.Next = n2
	n1 := &listNode{Val: 3, Next: n2}
	res := cycleByHash(n1)
	fmt.Println(res)
	res = cycleByFastLowPointer(n1)
	fmt.Println(res)
}
