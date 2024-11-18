package main

import (
	"fmt"
	"strings"
)

func rotateString(A, B string) bool {
	if strings.EqualFold(A, "") && strings.EqualFold(B, "") {
		return true
	}
	if len(A) != len(B) {
		return false
	}
	length := len(A)
	for i := 0; i < length; i++ {
		subFirst := A[0:1]
		subLast := A[1:]
		A = subLast + subFirst
		if strings.EqualFold(A, B) {
			return true
		}
	}
	return false
}
func main() {
	A := "abcde"
	B := "cdeab"
	res := rotateString(A, B)
	fmt.Println(res)
}
