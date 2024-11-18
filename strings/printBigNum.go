package main

import "fmt"

func printBigNum(n int) { //大数打印 非常巧妙，可以多参考
	if n < 1 {
		return
	}
	result := make([]byte, n)
	maxIndex := 1
	result[n-maxIndex] = '1'
loop:
	for {
		s := string(result[n-maxIndex:])
		fmt.Println(s)
		for i := 1; i <= maxIndex; i++ {
			if '9' != result[n-i] {
				result[n-i]++
				continue loop
			}
			result[n-i] = '0'
		}
		maxIndex++
		if maxIndex > n {
			break
		}
		result[n-maxIndex] = '1'
	}
}
func main() {
	printBigNum(3)
}
