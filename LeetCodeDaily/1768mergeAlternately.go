package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	n, m := len(word1), len(word2)
	ans := make([]byte, 0, n+m)
	for i := 0; i < n || i < m; i++ {
		if i < n {
			ans = append(ans, word1[i])
		}
		if i < m {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}

func main() {
	word1 := "abc"
	word2 := "pqr"
	res := mergeAlternately(word1, word2)
	fmt.Println(res)
}
