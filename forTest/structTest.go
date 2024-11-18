package main

import (
	"fmt"
	"os"
)

type ST1 struct {
	name string
	age  int
}
type ST2 struct {
	name string
	age  int
}

type ST struct {
	st1 ST1
	st2 ST2
}

func main() {
	s := &ST1{
		name: "",
	}
	if s != nil {
		fmt.Println(s.age)
	}
	var offset int64
	fmt.Println(offset)

	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")

}
