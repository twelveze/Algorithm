package main

import "fmt"

func minDistance( a string ,  b string ) int {
	lena := len(a)
	lenb := len(b)
	var dp = make([][]int, lena + 1)
	for i := range dp{
		dp[i] = make([]int, lenb + 1)
	}
	for i := 0; i <= lena; i++{
		for j := 0; j <= lenb; j++{
			if i == 0{
				dp[i][j] = j
			}else if j == 0{
				dp[i][j] = i
			}else if a[i - 1] == b[j - 1]{
				dp[i][j] = dp[i - 1][j - 1]
			}else{
				dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1
			}
		}
	}
	return dp[lena][lenb]
}
func min(a, b int) int{
	if a < b{
		return a
	}else{
		return b
	}
}
func main() {
	s := "horse"
	afters := "rouse"
	res := minDistance(s, afters)
	fmt.Println(res)
}
