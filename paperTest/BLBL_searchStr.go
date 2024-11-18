package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	var s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	res := strings.Index(s1, s2)
	fmt.Println(res)
}
