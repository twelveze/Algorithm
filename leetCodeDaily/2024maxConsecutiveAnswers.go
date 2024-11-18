package main

import (
	"Algorithm/Algorithmtest/LeetCodeDaily/common"
	"fmt"
)

func chooseByteToChange(answerKey string, k int, b byte) (res int) {
	left, limit := 0, 0
	for right := range answerKey {
		if answerKey[right] != b {
			limit++
		}
		if limit > k {
			if answerKey[left] != b {
				limit--
			}
			left++
		}
		res = common.Max(res, right-left+1)
	}
	return
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return common.Max(chooseByteToChange(answerKey, k, 'T'), chooseByteToChange(answerKey, k, 'F'))
}

func main() {
	answerKey := "TTFTTFTT"
	k := 1
	res := maxConsecutiveAnswers(answerKey, k)
	fmt.Println(res)
}
