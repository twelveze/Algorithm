package main

import "fmt"

func firstBadVersion(badVersion []bool) int{
	len := len(badVersion)
	left := 0
	right := len - 1
	for left < right{
		mid := (left + right) / 2
		if badVersion[mid]{
			left = mid + 1
		}else{
			right = mid
		}
	}
	return left
}
func main() {
	badVersion := []bool{true,true,true,true,false,false,false,false}
	res := firstBadVersion(badVersion)
	fmt.Println(res)
}
