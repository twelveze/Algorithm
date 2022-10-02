package main

import (
	"fmt"
	"io"
	"sort"
)

func totalPrice(category, price []int) []int{
	categoryLen := len(category)
	priceLen := len(price)
	min, max := 0, 0
	for i := 0; i < categoryLen; i++{
		min = min + category[categoryLen - i - 1] * price[i]
		max = max + category[categoryLen - i - 1] * price[priceLen - i - 1]
	}
	return []int{min,max}
}
func main() {
	for{
		var n, m int
		_,err := fmt.Scanf("%d %d",&n, &m)
		if err != nil{
			if err == io.EOF{
				break
			}
		}else{
			price := make([]int, n)
			category := make([]int, 0)
			categoryMap := make(map[string]int, 0)
			for i := 0; i < n; i++{
				fmt.Scan(&price[i])
			}
			for i := 0; i < m; i++{
				var s string
				fmt.Scanln(&s)
				categoryMap[s]++
			}
			for _, v := range categoryMap{
				category = append(category, v)
			}
			sort.Ints(category)
			sort.Ints(price)
			res := totalPrice(category, price)
			fmt.Println(res[0],res[1])
		}
	}
}
