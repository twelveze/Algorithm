package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
func main() {
	var n = 3
	res := bulbSwitch(n)
	fmt.Println(res)
}
