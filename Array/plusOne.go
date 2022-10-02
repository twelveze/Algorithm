package main

import "fmt"

func plusOne(digits []int) []int{
	var res []int
	carry := 0
	for i := len(digits) - 1; i >= 0; i--{
		digits[i] += carry //先加进位
		carry = 0 //进位变为0
		if i == len(digits) - 1{
			digits[i] += 1 //再加1
		}
		if digits[i] >= 10{
			carry = digits[i] / 10
			digits[i] = digits[i] % 10
		}
	}
	if carry != 0{
		res = make([]int, 1)
		res[0] = carry
		res = append(res, digits...)
	}else{
		res = digits
	}
	return res
}
func main() {
	digits := []int{9,9,9,9}
	res := plusOne(digits)
	fmt.Println(res)
}
