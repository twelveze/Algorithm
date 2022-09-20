package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("heihei")
	}()
	defer func() {
		fmt.Println("haha")
	}()
	panic("This is a panic")
}
