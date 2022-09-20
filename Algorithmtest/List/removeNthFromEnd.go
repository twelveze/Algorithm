package main

import "fmt"

//type listNode struct {
//	Val int
//	Next *listNode
//}
//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点
func removeNthFromEnd(head *listNode, n int) *listNode {
	result := &listNode{} //result为哨兵节点
	result.Next = head
	var pre *listNode
	cur := result
	i := 1
	for head != nil { //当head指向nil时，cur指向的就是要删除的节点
		if i >= n { //当head走了n步之后，cur也跟着走
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next //cur的pre指向cur.Next
	return result.Next
}
func main() {
	n5 := listNode{Val: 5, Next: nil}
	n4 := listNode{Val: 4, Next: &n5}
	n3 := listNode{Val: 3, Next: &n4}
	n2 := listNode{Val: 2, Next: &n3}
	n1 := listNode{Val: 1, Next: &n2}
	for tmp := removeNthFromEnd(&n1, 2); tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Val, " ")
	}
}
