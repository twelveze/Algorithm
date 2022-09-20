package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string{
	substr := strs[0]
	for _, s := range strs{
		for strings.Index(s, substr) == -1{
			substr = substr[:len(substr) - 1]
		}
	}
	return substr
}
func main() {
	strs := []string{"flow","flower","flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}