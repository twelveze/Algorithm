package main

import "fmt"

func Add(i int) {
	i += 1
	fmt.Println(i)
}
func AddByPointer(i *int) {
	*i += 1
	fmt.Println(*i)
}
func main() {
	i := 1
	Add(i)
	fmt.Println(i)
	AddByPointer(&i)
	fmt.Println(i)
}
