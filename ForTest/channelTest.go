package main

import "fmt"

func main() {

	testChannel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		testChannel <- i
	}

	for v := range testChannel { //所以说for range也是会把channel的元素弹出的
		fmt.Println(v)
	}

	fmt.Println(<-testChannel)
}