package main

import "fmt"

func main() {
	mum := 1
	child := 2
	arr := make([]float64, 0)
	arr = append(arr, float64(child) / float64(mum))
	for i := 1; i < 20; i++{
		temp := child
		child = child + mum
		mum = temp
		arr = append(arr, float64(child) / float64(mum))
	}
	var sum float64
	for i := 0; i < len(arr); i++{
		sum = sum + arr[i]
	}
	fmt.Printf("%.1f", sum)
}
