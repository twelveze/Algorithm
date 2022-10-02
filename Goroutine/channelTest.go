package main

import "fmt"

func main() {
	nochan := make(chan int)
	go func(ch chan int) {
		ch <- 100
		fmt.Println("send data", 100)
		close(ch)
		fmt.Println("goroutine exit")
	}(nochan)
	data := <-nochan
	fmt.Println("receive data is ", data)
	data, ok := <-nochan
	if !ok {
		fmt.Println("receive close chan")
		fmt.Println("receive data is ", data)
	}
	fmt.Println("main exited")
}
