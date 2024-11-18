package main

import (
	"fmt"
	"sort"
)

func originalDigits(s string) string {
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	cnt := [10]int{}
	cnt[0] = count['z']
	cnt[2] = count['w']
	cnt[4] = count['u']
	cnt[6] = count['x']
	cnt[8] = count['g']

	cnt[3] = count['h'] - cnt[8]
	cnt[5] = count['f'] - cnt[4]
	cnt[7] = count['s'] - cnt[6]

	cnt[1] = count['o'] - cnt[0] - cnt[2] - cnt[4]
	cnt[9] = count['i'] - cnt[5] - cnt[6] - cnt[8]

	res := make([]int, 0)
	for i := 0; i < 10; i++ {
		for cnt[i] > 0 {
			res = append(res, i)
			cnt[i]--
		}
	}
	sort.Ints(res)
	var resString []byte
	for i := 0; i < len(res); i++ {
		resString = append(resString, byte(res[i]+'0'))
	}
	return string(resString)
}

func main() {
	var s = "fviefuro"
	res := originalDigits(s)
	fmt.Println(res)
}
