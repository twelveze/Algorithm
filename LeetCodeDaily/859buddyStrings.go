package main

import "fmt"

// 1、长度要一致；2、字符串要么有两个字符不相同且可以互换；3、要么字符串相同且最大频次字母大于1
func buddyStrings(s, goal string) bool{
	if len(s) != len(goal){
		return false
	}
	if s == goal{ // 字符串相同且最大频次字母大于1
		trace := make(map[rune]int)
		for  _, b := range s{
			trace[b]++
			if trace[b] > 1{
				return true
			}
		}
		return false
	}

	first, second := -1, -1
	for i := range s{ //有且只有两个字符串不同
		if s[i] != goal[i]{
			if first == -1{
				first = i
			}else if second == -1{
				second = i
			}else{
				return false
			}
		}
	}
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}

func main() {
	var s = "aaaaaaabc"
	var goal = "aaaaaaacb"
	res := buddyStrings(s, goal)
	fmt.Println(res)
}
