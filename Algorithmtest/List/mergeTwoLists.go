package main

import (
	"fmt"
)
//type listNode struct {
//	Val int
//	Next *listNode
//}
func mergeTwoLists(l1 *listNode, l2 *listNode) *listNode{
	head := &listNode{}
	result := head
	for l1 != nil && l2 != nil{
		if l1.Val > l2.Val{
			head.Next = l2
			l2 = l2.Next
		}else{
			head.Next = l1
			l1 = l1.Next
		}
		head = head.Next
	}
	if l1 != nil{
		head.Next = l1
	}
	if l2 != nil{
		head.Next = l2
	}
	return result.Next
}
func main() {
	n3 := listNode{Val: 4, Next: nil}
	n2 := listNode{Val: 2, Next: &n3}
	n1 := listNode{Val: 1, Next: &n2}

	n6 := listNode{4,nil}
	n5 := listNode{3,&n6}
	n4 := listNode{1,&n5}
	for res := mergeTwoLists(&n1, &n4); res != nil; res = res.Next {
		fmt.Print(res.Val, " ")
	}
}