package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32 = 1
	var b int32 = 2
	atomic.CompareAndSwapInt32(&a, a, b)
	fmt.Println(a, b)

	var m atomic.Value
	m.Store(map[int]string{
		2: "hahaha",
		3: "heiheihei",
	})
	fmt.Println(m.Load())
}
