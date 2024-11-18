package main

import "fmt"

func test(a int) (i int) {
	defer func() {
		i++
	}()
	i = i + a
	return
}
func main() {
	i := test(1)
	fmt.Println(i)
}
