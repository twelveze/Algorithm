package main

import (
	"fmt"
	"io"
)

func check(messages []int, n int, method [][]int) (res []int) {
	for i := 0; i < n; i++ {
		ans := 0
		for len(messages) > 0 {
			index := 0
			count := 0
			for j := 0; j < method[i][0]; j++ {
				if count == method[i][1] {
					break
				} else {
					index++
				}
				if messages[j] == 1 {
					count++
				}
			}
			lanjie := messages[index : index+method[i][2]]
			for k := 0; k < len(lanjie); k++ {
				if lanjie[k] == 1 {
					ans++
				}
			}
			messages = messages[index+method[i][2]:]
		}
		res = append(res, ans)
	}
	return
}
func main() {
	var message int
	messages := make([]int, 0)
	for {
		_, err := fmt.Scan(&message)
		messages = append(messages, message)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	var n int
	fmt.Scanf("%d", &n)
	method := make([]int, 0)
	methods := make([][]int, 0)
	for i := 0; i < n; i++ {
		var x, y, z int
		fmt.Scanf("%d %d %d", &x, &y, &z)
		method = append(method, x, y, z)
		methods = append(methods, method)
	}
	res := check(messages, n, methods)
	fmt.Println(res)
}
