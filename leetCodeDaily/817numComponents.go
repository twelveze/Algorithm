package main

import (
	"fmt"
	"model"
)

func numComponents(head *model.ListNode, nums []int) int {
	ans := 0
	numMap := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		numMap[v] = struct{}{}
	}
	for flag := true; head != nil; head = head.Next {
		if _, ok := numMap[head.Val]; !ok { //如果不存在,flag置为true,这样当下次出现存在的数字时就会+1
			flag = true
		} else if flag {
			flag = false //连续出现存在的数字只会被记为1
			ans++
		}
	}
	return ans
}
func main() {
	n4 := &model.ListNode{Val: 3, Next: nil}
	n3 := &model.ListNode{Val: 2, Next: n4}
	n2 := &model.ListNode{Val: 1, Next: n3}
	n1 := &model.ListNode{Val: 0, Next: n2}
	nums := []int{0, 3, 1}
	res := numComponents(n1, nums)
	fmt.Println(res)
}
