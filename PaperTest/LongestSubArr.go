package main

import "fmt"
func MAX(a, b int) int{
	if a > b {
		return a
	}else{
		return b
	}
}
func longestSubArr(str1, str2 string) int{
	res := make([][]int, len(str1) + 1)
	for i, _ := range res{
		res[i] = make([]int, len(str2) + 1)
	}
	for i := 1; i <= len(str1); i++{
		for j := 1; j <= len(str2); j++{
			if str1[i - 1] == str2[j - 1]{
				res[i][j] = res[i - 1][j - 1] + 1
			}else{
				res[i][j] = MAX(res[i - 1][j], res[i][j - 1])
			}
		}
	}
	return res[len(str1)][len(str2)]
}
func main()  {
	str1 := "1A2C3D4B56"
	str2 := "B1D23A456A"
	res := longestSubArr(str1, str2)
	fmt.Println(res)
}
