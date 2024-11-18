package main

import (
	"container/ring"
	"fmt"
)

func main() {

	r := ring.New(10)

	//给闭环中的元素赋值
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	//循环打印闭环中的元素值
	r.Do(
		func(p interface{}) {
			println(p.(int))
		})

	//获得当前元素之后的第5个元素
	r5 := r.Move(5)
	fmt.Println(r.Value)
	fmt.Println(r5.Value)
	fmt.Println("xxx")
	//链接当前元素r与r5，相当于删除了r与r5之间的元素
	r1 := r.Link(r5)
	fmt.Printf("r1 = %d\n", r1.Value)
	r5.Do(
		func(p interface{}) {
			println(p.(int))
		})
}
