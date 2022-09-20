package main
import (
	"fmt"
	"time"
)

func printer()  {
	defer func() {
		if err := recover(); err!=nil{
			fmt.Println("test发生错误",err)
		}
	}()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
	//定义一个map的错误用法导致error
	var myMap map[int]string
	myMap[0] = "go"
	fmt.Println("crash")
}

func main() {

	go printer()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=",i)
		time.Sleep(time.Second)
	}
}