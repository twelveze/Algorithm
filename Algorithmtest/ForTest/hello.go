package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("hello")
	str := "what happen?"
	fmt.Println(str)
	fmt.Println(runtime.NumCPU())

	var nanoTime int64 = 1664100842163119000
	timer := time.Unix(0, nanoTime)
	fmt.Println(timer)
}
