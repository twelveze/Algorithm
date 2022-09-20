package main

import "fmt"

func main() {
	s1 := []string{"hello", "世界", "gaoxiao"}
	for i := 0; i < len(s1); i++ {
		fmt.Println(len([]byte(s1[i])))
	}
	fmt.Println()
	for i := 0; i < len(s1); i++ {
		fmt.Println(len([]rune(s1[i])))
	}
}
