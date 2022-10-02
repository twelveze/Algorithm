package main

import "fmt"

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
func CombinationsOfNumbers(digits string) []string{
	if len(digits) == 0{
		return []string{}
	}
	var res []string
	backtrack(&res, 0, digits, "")
	return res
}
func backtrack(res *[]string, index int, digits string, combination string){
	if index == len(digits){
		*res = append(*res, combination)
		return
	}
	phoneNum := string(digits[index])
	phoneStr := phoneMap[phoneNum]
	for i := 0; i < len(phoneStr); i++{
		backtrack(res, index + 1, digits, combination + string(phoneStr[i]))
	}
}
func main() {
	digits := "22"
	res := CombinationsOfNumbers(digits)
	fmt.Println(res)
}
