package main

import "fmt"

func f(x int) func(int) int {
	g := func(y int) int {
		x++
		return x + y
	}
	// 返回闭包
	return g
}

func main() {
	// 将函数的返回结果"闭包"赋值给变量a
	a := f(3)
	// 调用存储在变量中的闭包函数
	res := a(5)
	fmt.Println(res)
	res = a(8)
	fmt.Println(res)
}
