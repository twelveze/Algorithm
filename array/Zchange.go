package main

import (
	"fmt"
	"strings"
)

func zChange(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	flag := -1
	row := 0
	for i := 0; i < len(s); i++ {
		rows[row] += string(s[i])
		if i%(numRows-1) == 0 { //在第一行和最后一行的时候变换方向
			flag = -flag
		}
		row += flag
	}
	return strings.Join(rows, "")
}
func main() {
	s := "LEETCODEISHIRING"
	ss := zChange(s, 4)
	fmt.Println(ss)
}
