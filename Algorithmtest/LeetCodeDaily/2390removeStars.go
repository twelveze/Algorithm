package main

import "fmt"

func removeStars(s string) string {
	remove := make([]bool, len(s)) //第i位是否要删除
	count := 0                     //需要删除的个数
	//从右往左遍历，看左边需要删除的字符
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			count++
			remove[i] = true
		} else {
			if count != 0 {
				remove[i] = true
				count--
			}
		}
	}
	ans := []byte{}
	for i := 0; i < len(s); i++ {
		if !remove[i] { //没有被移除的才加入到ans
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
func main() {
	str := "leet**cod*e"
	res := removeStars(str)
	fmt.Println(res)
}
