package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func addTwoNumbers(l1, l2 *listNode) *listNode {
	list := &listNode{}
	result := list
	temp := 0
	for l1 != nil || l2 != nil || temp != 0 {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		list.Next = &listNode{Val: temp % 10, Next: nil}
		list = list.Next
		temp = temp / 10
	}
	return result.Next
}
func main() {
	l1n3 := &listNode{Val: 3, Next: nil}
	l1n2 := &listNode{Val: 4, Next: l1n3}
	l1n1 := &listNode{Val: 2, Next: l1n2}

	l2n3 := &listNode{Val: 4, Next: nil}
	l2n2 := &listNode{Val: 6, Next: l2n3}
	l2n1 := &listNode{Val: 5, Next: l2n2}
	for res := addTwoNumbers(l1n1, l2n1); res != nil; res = res.Next {
		fmt.Println(res.Val)
	}
}
