package main

import "fmt"
func NoSame(phone, web []int)int{
	phonemap := make(map[int]int,0)
	res := 0
	for i := 0; i < len(phone); i++{
		if _, ok := phonemap[phone[i]]; ok{
			phonemap[phone[i]]++
		}else{
			phonemap[phone[i]] = 1
		}
	}
	for i := 0; i < len(web); i++{
		if _, ok := phonemap[web[i]]; !ok{
			res++
		}else{
			phonemap[web[i]] = 0
		}
	}
	for _, v := range phonemap{
		res = res + v
	}
	return res
}
func main() {
	var n, m int
	fmt.Scanf("%d %d",&n,&m)
	phone := make([]int, 0)
	web := make([]int, 0)
	for i := 0; i < n; i++{
		var temp int
		fmt.Scan(&temp)
		phone = append(phone, temp)
	}
	for i := 0; i < m; i++{
		var temp int
		fmt.Scan(&temp)
		web = append(web, temp)
	}
	res := NoSame(phone, web)
	fmt.Println(res)
}
