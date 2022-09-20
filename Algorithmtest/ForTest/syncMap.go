package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("address", map[string]string{
		"province": "江苏",
		"city":     "南京",
	})
	v, _ := m.Load("address")
	fmt.Println(v)

	map1 := map[string]string{"haha": "he"}
	fmt.Println(map1["haha"])
}
