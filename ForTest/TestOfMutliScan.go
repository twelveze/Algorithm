package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := make([]int,0)
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	str := strings.Replace(input, "\n", "", -1)
	str = strings.Replace(str, " ", "", -1)
	for i := 0; i < len(str); i++{
		temp, _ := strconv.Atoi(string(str[i]))
		arr = append(arr, temp)
	}
	fmt.Println(arr)
}
