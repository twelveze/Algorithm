package main

import "fmt"

func stairs(n int, bads map[int]bool)int{
	dp := make([]int, n + 1)
	for i := 0; i < n + 1;i++{
		dp[i] = 0
	}
	dp[0] = 1
	for i := 1; i <= n; i++{
		if _, ok := bads[i]; ok{
			continue
		}else{
			temp := 0
			if _, ok := bads[i - 1]; !ok && i - 1 >= 0{
				temp += dp[i - 1]
			}
			if _, ok := bads[i - 2]; !ok && i - 2 >= 0{
				temp += dp[i - 2]
			}
			dp[i] = temp
		}
	}
	return dp[n]
}

func main() {
	var n , bad int
	fmt.Scanf("%d %d", &n, &bad)
	bads := make(map[int]bool, bad)
	for i := 0; i < bad; i++{
		var badStairs int
		fmt.Scan(&badStairs)
		bads[badStairs] = true
	}
	res := stairs(n, bads)
	fmt.Println(res)
}
