package main

import (
	"fmt"
	"reflect"
	"time"
)

type s struct {
	i int
}

func main() {

	//var a interface{} = &s{i: 2}
	//var t interface{} = time.Time{}
	//可以通过这张方式把数据库的数据读到[]interface里
	var dest []interface{}
	dest = append(dest, new(time.Time))

	switch v := dest[0].(type) {
	case *s, s:
		fmt.Println("yes")
	case *time.Time, time.Time:
		fmt.Println("time yes")
		fmt.Println(v.(*time.Time))
	default:
		fmt.Println(v.(*s))
		fmt.Println(v)
	}
	var i interface{}
	i = s{
		i: 2,
	}
	value := reflect.ValueOf(i)
	kind := value.Kind()
	types := value.Type()
	fmt.Println(kind, types)
}
