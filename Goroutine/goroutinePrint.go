package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan struct{}, wg *sync.WaitGroup) {
		for i := 1; i <= 10; i++ {
			ch <- struct{}{}
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan struct{}, wg *sync.WaitGroup) {
		for i := 1; i <= 10; i++ {
			<-ch
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
