package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func goModel(ctx context.Context){
	fmt.Println("come in goModel")
	defer fmt.Println("go out of goModel")
	defer wg.Done()
	LOOP:
	for{
		fmt.Println("model go say hello")
		time.Sleep(time.Second*1)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}
func main() {
	fmt.Println("come in main()")
	defer fmt.Println("out of main()")
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go goModel(ctx)
	time.Sleep(time.Second*5)
	cancel()
	wg.Wait()
}
