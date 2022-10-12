package main

import (
	"Algorithm/Model"
	"fmt"
)

func numComponents(head *Model.ListNode, nums []int) int {
	ans := 0
	numMap := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		numMap[v] = struct{}{}
	}
	for flag := false; head != nil; head = head.Next {
		if _, ok := numMap[head.Val]; !ok {
			flag = false
		} else if !flag {
			flag = true
			ans++
		}
	}
	return ans
}
func main() {
	n4 := &Model.ListNode{Val: 3, Next: nil}
	n3 := &Model.ListNode{Val: 2, Next: n4}
	n2 := &Model.ListNode{Val: 1, Next: n3}
	n1 := &Model.ListNode{Val: 0, Next: n2}
	nums := []int{0, 3, 1}
	res := numComponents(n1, nums)
	fmt.Println(res)
}
