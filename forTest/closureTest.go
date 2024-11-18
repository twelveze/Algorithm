package main

import "fmt"

func a() int {
	var i int
	defer add(i)
	i += 100
	return i
}

//b中实际上是通过一个闭包调用，将i的指针传递给闭包，闭包读取值拷贝给add，这个时候add是改变不了i的值的
func b() int {
	var i int
	defer func() {
		add(i)
	}()
	i += 100
	return i
}
func c() int {
	i := 100
	defer func() {
		i += 100
	}()
	return i
}
func d() (i int) {
	i = 100
	defer func() {
		i += 1
	}()
	return
}

func add(i int) {
	i += 1
	fmt.Println(i)
}
func main() {
	res := a()
	//1
	//100
	fmt.Println(res)
	res = b()
	//101
	//100
	fmt.Println(res)
	res = c()
	//100
	fmt.Println(res)
	res = d()
	//101
	fmt.Println(res)
}
