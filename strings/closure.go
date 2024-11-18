package main

import "fmt"

func main() {
	n := 0
	f := func() int {
		n += 1
		return n
	}

	fmt.Println(f())
	fmt.Println(f())

	fmt.Println(n)
}
