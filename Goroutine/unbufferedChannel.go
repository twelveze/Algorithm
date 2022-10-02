package main

import (
	"fmt"
)

// 申请一个无缓冲的双向int型chan
var c chan int = make(chan int)

func foo() {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("foo协程发送的数据是：", i)
	}
}
func main() {
	go foo()
	//for i := 0; i < 10; i++ {
	//	fmt.Println("主协程读到的数据是：", <-c)
	//}
	for v := range c {
		fmt.Println("主协程读到的数据是：", v)
	}
}
