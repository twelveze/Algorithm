package main

import (
	"fmt"
)

type demoStruct struct {
	name string
}
type demo struct {
	name string
}

func xxx(x int) (res int) {
	defer func() {
		res = res + 1
	}()
	return 2
}

type people struct{}
type teacher struct {
	people
}
type student struct {
	Name string
	Age  int
}

func parseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for _, v := range m {
		fmt.Println(v.Name)
	}
}
func (p *people) sayHello() {
	fmt.Println("hello")
	p.work()

}
func (p *people) work() {
	fmt.Println("people work")
}
func (t *teacher) work() {
	fmt.Println("teacher work")
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	//var num int64 = -7
	//fmt.Println(num >> 1)
	//fmt.Println(strconv.IntSize) //int位数

	//a := &demoStruct{
	//	name: "halo",
	//}
	//b := &demo{
	//	name: "halo",
	//}
	//if reflect.DeepEqual(a, b) {
	//	fmt.Println("yes")
	//} else {
	//	fmt.Println("no")
	//}

	//a := [3]int{1, 2, 3}
	//slice := a[:2]
	//slice = append(slice, []int{0, 0, 0}...)
	//slice[0] = 4
	//fmt.Println(a)
	//fmt.Println(slice)

	//res := xxx(2)
	//fmt.Println(res)

	//slice := make([]int, 3, 3)
	//slice = append(slice, []int{1, 2, 3}...)
	//fmt.Println(slice)

	//t := &teacher{
	//	people{},
	//}
	//t.sayHello()
	//t.work()

	//parseStudent()

	//runtime.GOMAXPROCS(4)
	//wg := sync.WaitGroup{}
	//wg.Add(20)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println("first i: ", i)
	//		wg.Done()
	//	}()
	//}
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("second i: ", i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()

	//a := 1
	//b := 2
	//defer calc("1", a, calc("10", a, b))
	//a = 0
	//defer calc("2", a, calc("20", a, b))
	//b = 1

	//var peo People = &Stduent{}
	//think := "bitch"
	//fmt.Println(peo.Speak(think))

	//n := runtime.GOMAXPROCS(1) //将cpu设置为1核
	//fmt.Println(n)
	//wg := sync.WaitGroup{}
	//wg.Add(10)
	////将cpu设置为1核时，下面两个进程将会互相争夺，谁抢到，谁就会开始不断打印
	////所以打印结果是大片的1或0(谁抢到打印谁)，下面紧接着打印另外一个数
	////相比较给予比进程数多的核来讲，打印的数更加密集，因为多个核，两个进程竞争变小
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Print("goroutine: ", i) //子go程
	//		wg.Done()
	//	}()
	//
	//	fmt.Print(1) //主go程
	//}
	//wg.Wait()

}
