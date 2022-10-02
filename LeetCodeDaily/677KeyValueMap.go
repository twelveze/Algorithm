package main

import "fmt"

type MapSum struct {
	m map[string]int
	pre map[string]int
}


func Constructor() MapSum {
	return MapSum{map[string]int{}, map[string]int{}}
}


func (m *MapSum) Insert(key string, val int)  {
	delta := val
	if m.m[key] > 0 {	//如果已经存在，就计算出与先前值的差，然后更新所有pre
		delta -= m.m[key]
	}
	m.m[key] = val
	for i := range key {
		m.pre[key[:i+1]] += delta
	}
}


func (m *MapSum) Sum(prefix string) int {
	return m.pre[prefix]
}
func main() {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	res := mapSum.Sum("ap")
	fmt.Println(res)
	mapSum.Insert("app", 2)
	res = mapSum.Sum("ap")
	fmt.Println(res)
}
