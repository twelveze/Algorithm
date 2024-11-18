package main

import (
	"fmt"
)

var exit = make(chan bool)
var i = 0

func printA(chanA, chanB chan bool) {
	for {
		if ok := <-chanA; ok {
			fmt.Println("A")
			chanB <- true
		}
	}
}
func printB(chanB, chanC chan bool) {
	for {
		if ok := <-chanB; ok {
			fmt.Println("B")
			chanC <- true
		}
	}
}
func printC(chanC, chanA, exit chan bool) {
	for {
		if ok := <-chanC; ok {
			fmt.Println("C")
			chanA <- true
			i++
		}
		if i == 2 {
			exit <- true
		}
	}

}
func main() {
	chanA := make(chan bool, 1)
	chanB := make(chan bool, 1)
	chanC := make(chan bool, 1)

	go printA(chanA, chanB)
	go printB(chanB, chanC)
	go printC(chanC, chanA, exit)
	chanA <- true
	<-exit
}
