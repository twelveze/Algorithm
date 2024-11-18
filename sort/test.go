package main

import "fmt"

func sum(a int, b int) (s int, i int) {

	defer func(n int) {
		a--
		fmt.Println("2:b-a=", b, a, n-a)
		i++
		s += i
	}(b)
	defer func(n int) {
		b--
		fmt.Println("1:a-b=", a, b, n-b)
		i++
		s = n + b
	}(a)
	return s, i
}
func main() {
	s, i := sum(10, 10)
	fmt.Print("3:s=", s, ",i=", i)
}
