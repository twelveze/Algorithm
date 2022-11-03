package main

import (
	"fmt"
	"strings"
)

//看看重复k次后是不是sequence的子串
func maxRepeating(sequence string, word string) int {
	for k := len(sequence) / len(word); k > 0; k-- {
		if strings.Contains(sequence, strings.Repeat(word, k)) {
			return k
		}
	}
	return 0
}

func main() {
	sequence := "ababc"
	word := "ab"
	res := maxRepeating(sequence, word)
	fmt.Println(res)
}
