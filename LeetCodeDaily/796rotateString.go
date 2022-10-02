package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool{
	if len(s) != len(goal){
		return false
	}
	if strings.Contains(s+s, goal){
		return true
	}else{
		return false
	}
}

func main() {
	s1 := "abcde"
	s2 := "abced"
	res := rotateString(s1, s2)
	if res{
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}
