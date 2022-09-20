package main

import (
	"Algorithm/Algorithmtest/Model"
	"fmt"
)

func reverseList1(head *Model.ListNode) *Model.ListNode {
	var prev *Model.ListNode //pre节点代表的是当前节点的前一个节点
	cur := head              // cur代表当前节点
	for cur != nil {
		temp := cur.Next // temp节点代表当前节点cur的下一个节点，因为cur与cur.Next要先断开，为防止找不到，用temp保存
		cur.Next = prev  //当前节点指向它的前一个节点
		prev = cur       //当前节点变成下一个节点的prev节点
		cur = temp       //当前节点变成下一个节点
	}
	return prev
}

//处理后的拼接顺序为head->start->reverseEnd->reverseStart->end
func reverseBetween(head *Model.ListNode, left int, right int) *Model.ListNode {
	//reverseStart是待翻转链表的头节点
	var reverseStart *Model.ListNode
	//reverseEnd是待反转链表的尾节点
	var reverseEnd *Model.ListNode
	//start是待翻转链表的头节点的pre节点，这里初始化的时候必须赋个值，不然报错
	var dummyNode = &Model.ListNode{Val: -1, Next: nil}
	//end是待翻转链表的尾节点的next节点
	var end *Model.ListNode
	dummyNode.Next = head
	start := dummyNode
	for i := 0; i < left-1; i++ {
		start = start.Next
	}
	reverseStart = start.Next

	reverseEnd = start
	for i := 0; i < right-left+1; i++ {
		reverseEnd = reverseEnd.Next
	}
	//暂存一下end的next节点，后续要断开
	end = reverseEnd.Next
	//把中间要翻转的链表断开
	start.Next = nil
	reverseEnd.Next = nil

	//翻转
	reverseList1(reverseStart)

	//拼接
	start.Next = reverseEnd
	reverseStart.Next = end

	return dummyNode.Next
}
func main() {
	//n5 := &Model.ListNode{Val: 5, Next: nil}
	//n4 := &Model.ListNode{Val: 4, Next: n5}
	//n3 := &Model.ListNode{Val: 3, Next: n4}
	n2 := &Model.ListNode{Val: 5, Next: nil}
	n1 := &Model.ListNode{Val: 3, Next: n2}
	root := n1
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
	root = reverseBetween(n1, 1, 2)
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
}
