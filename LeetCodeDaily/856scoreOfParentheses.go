package main

import (
	"Algorithm/Model"
	"fmt"
)

//栈思想
func scoreOfParentheses(s string) int {
	stack := []int{0}
	for _, c := range s {
		if c == '(' {
			stack = append(stack, 0)
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += Model.Max(2*v, 1)
		}
	}
	return stack[0]
}
func main() {
	s := "(()(()))"
	res := scoreOfParentheses(s)
	fmt.Println(res)
}
