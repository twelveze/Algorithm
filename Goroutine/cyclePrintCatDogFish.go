package main

import (
	"fmt"
	"sync"
)

func printCat(counter int, catch, dogch chan struct{}, wg *sync.WaitGroup){
	for{
		if counter >= 10{
			wg.Done()
			return
		}
		<- catch
		fmt.Println("cat")
		counter++
		dogch <- struct{}{}
	}
}
func printDog(counter int, dogch, fishch chan struct{}, wg *sync.WaitGroup){
	for{
		if counter >= 10{
			wg.Done()
			return
		}
		<- dogch
		fmt.Println("dog")
		counter++
		fishch <- struct{}{}
	}
}
func printFish(counter int, fishch, catch chan struct{}, wg *sync.WaitGroup){
	for{
		if counter >= 10{
			wg.Done()
			return
		}
		<- fishch
		counter++
		fmt.Println("fish")
		catch <- struct{}{}
	}
}
func main() {
	wg := &sync.WaitGroup{}
	catch := make(chan struct{},1)
	dogch := make(chan struct{},1)
	fishch := make(chan struct{},1)
	var catcounter int
	var dogcounter int
	var fishcounter int
	wg.Add(3)

	go printCat(catcounter, catch, dogch, wg)
	go printDog(dogcounter, dogch, fishch, wg)
	go printFish(fishcounter, fishch, catch, wg)

	catch <- struct{}{}
	wg.Wait()
}