package main

import "fmt"

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	if k <= (1 << (n - 2)) {
		return kthGrammar(n-1, k)
	}
	return kthGrammar(n-1, k-(1<<(n-2))) ^ 1
}

func main() {
	n := 2
	k := 2
	res := kthGrammar(n, k)
	fmt.Println(res)
}
