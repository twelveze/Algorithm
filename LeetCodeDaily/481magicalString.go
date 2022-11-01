package main

import "fmt"

func magicalString(n int) int {
	if n < 4 {
		return 1
	}
	i, j := 2, 3 //i代表现有的组数,j代表字符串的长度(i=2,j=3意味着字符串从"122"开始)
	res := 1
	s := []byte("122")
	for j < n {
		size := s[i] - '0'         //下一个元素应该有的个数
		elem := 3 - (s[j-1] - '0') //元素应该是1 or 2
		for size > 0 && j < n {
			s = append(s, '0'+elem)
			if elem == 1 {
				res++
			}
			j++
			size--
		}
		i++
	}
	return res
}

func main() {
	n := 6
	res := magicalString(n)
	fmt.Println(res)
}
