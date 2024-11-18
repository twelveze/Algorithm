package main

import (
	"fmt"
	"model"
)

func reverseList(head *model.ListNode) *model.ListNode {
	var prev *model.ListNode //pre节点代表的是当前节点的前一个节点
	cur := head              // cur代表当前节点
	for cur != nil {
		temp := cur.Next // temp节点代表当前节点cur的下一个节点，因为cur与cur.Next要先断开，为防止找不到，用temp保存
		cur.Next = prev  //当前节点指向它的前一个节点
		prev = cur       //当前节点变成下一个节点的prev节点
		cur = temp       //当前节点变成下一个节点
	}
	return prev
}
func main() {
	n3 := &model.ListNode{Val: 3, Next: nil}
	n2 := &model.ListNode{Val: 2, Next: n3}
	n1 := &model.ListNode{Val: 1, Next: n2}
	root := n1
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
	root = reverseList(n1)
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
}
