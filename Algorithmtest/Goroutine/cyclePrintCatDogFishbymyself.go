package main

import (
	"fmt"
	"sync"
)

func Cat(catch, dogch chan struct{}, wg *sync.WaitGroup){
	for i := 0; i <= 10; i++{
		<-catch
		fmt.Println("cat",i)
		dogch <- struct{}{}
	}
	wg.Done()
}
func Dog(dogch, fishch chan struct{}, wg *sync.WaitGroup){
	for i := 0; i <= 10; i++{
		<-dogch
		fmt.Println("dog",i)
		fishch <- struct{}{}
	}
	wg.Done()
}
func Fish(fishch, catch chan struct{}, wg *sync.WaitGroup){
	for i := 0; i <= 10; i++{
		<-fishch
		fmt.Println("fish",i)
		catch <- struct{}{}
	}
	wg.Done()
}
func main() {
	wg := &sync.WaitGroup{}
	catch := make(chan struct{},1)
	dogch := make(chan struct{},1)
	fishch := make(chan struct{},1)

	wg.Add(3)

	go Cat(catch, dogch, wg)
	go Dog(dogch, fishch, wg)
	go Fish(fishch, catch, wg)

	catch <- struct{}{}
	wg.Wait()
}