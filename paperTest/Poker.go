package main

import (
	"fmt"
	"sort"
)

func Poker(poker []string) (res []string) {
	sort.Strings(poker)
	return poker
}
func main() {
	var poker []string
	for i := 0; i < 7; i++ {
		var str string
		fmt.Scan(&str)
		poker = append(poker, str)
	}
	res := Poker(poker)
	fmt.Println(res)
}
