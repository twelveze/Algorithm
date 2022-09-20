package main

import "fmt"

func main() {
	str := "你好啊"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))
}
