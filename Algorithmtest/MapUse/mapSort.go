package main

import (
	"fmt"
	"sort"
)

func main() {
	var m = map[string]int{
		"hello":         0,
		"morning":       1,
		"keke":          2,
		"jame":   		 3,
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}