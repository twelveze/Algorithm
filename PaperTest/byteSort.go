package main

import (
	"fmt"
	"sort"
)

func byteSort(res []string, str *string, start, end int) []string{
	strByte := []byte(*str)
	if end <= 1{
		return []string{*str}
	}
	if start == end{
		res = append(res, string(strByte))
		return res
	}else{
		for i := start; i <= end; i++{
			strByte[i], strByte[start] = strByte[start], strByte[i]
			*str = string(strByte)
			res = byteSort(res, str, start + 1, end)
		}
	}
	return res
}
func main() {
	var str string
	fmt.Scanln(&str)
	var res []string
	res = byteSort(res, &str, 0, len(str) - 1)
	sort.Strings(res)
	for i := 0; i < len(res); i++{
		fmt.Println(res[i])
	}
}
