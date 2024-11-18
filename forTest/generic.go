package main

import (
	"encoding/json"
	"fmt"
)

// 具体的struct
type A struct {
	Id int64
}
type B struct {
	Id int32
}

// 泛型接口约束，struct的集合
type modelTest interface {
	A | B
}

// 第三方接口返回的数据
type ThirdResult[T modelTest] struct {
	Code    int
	Message string
	Result  []T //这里是 struct 的切片
}

// 解析数据到结构体
func parseData[T modelTest](data []byte) (items ThirdResult[T]) {
	result := []T{}
	thirdRes := ThirdResult[T]{Result: result}
	err := json.Unmarshal(data, &thirdRes)
	if err != nil {
		return ThirdResult[T]{}
	}
	return thirdRes
}

//泛型
func main() {
	str := "{\"code\":200, \"message\":\"Success\",\"result\": [{\"id\":1},{\"id\":2}]}" //json转义成string
	items := parseData[A]([]byte(str))
	fmt.Println(items)
}
