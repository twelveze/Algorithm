package main

import (
	"fmt"
)

//func maxProduct(words []string) int{
//	res := 0
//	for i := 0; i < len(words); i++{
//		for j := i + 1; j < len(words); j++{
//			if check(words[i], words[j]){
//				if len(words[i]) * len(words[j]) > res{
//					res =len(words[i]) * len(words[j])
//				}
//			}
//		}
//	}
//	return res
//}
//func check(word1, word2 string) bool{
//	stringMap := make(map[byte]bool)
//	for i := 0; i < len(word1); i++{
//		stringMap[word1[i]] = true
//	}
//	for i := 0; i < len(word2); i++{
//		if stringMap[word2[i]]{
//			return false
//		}
//	}
//	return true
//}
func maxProduct(words []string) (ans int) { //通过位运算法来快速判断两个单词有没有重复字母
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}

	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}
func main() {
	var words = []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	res := maxProduct(words)
	fmt.Println(res)
}
