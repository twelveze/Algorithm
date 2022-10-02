package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	nochan := make(chan int)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5
		fmt.Println("send data: ", 5)
		wg.Done()
	}(nochan, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		data := <-ch
		fmt.Println("receive data:",data)
		wg.Done()
	}(nochan,wg)
	wg.Wait()
}
