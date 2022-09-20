package main

import (
	"fmt"
	"sync"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	w := sync.WaitGroup{}
	w.Add(100)
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
			w.Done()
		}
		close(naturals)
	}()
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	//这里需要注意：wait和close，必须是基于sizes的循环的并发
	//go func() {
	// w.Wait()
	//}()
	//close(naturals)
	for x := range squares {
		fmt.Println(x)
	}
	w.Wait()
}
