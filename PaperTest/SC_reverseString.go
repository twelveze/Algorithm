package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) (res string){
	temp := make([]byte, 0)
	for i := 0; i < len(s); i++{
		if s[i] == ' '{
			if i >= 1 && s[i - 1] == ' '{
				continue
			}
			for j := 0; j < len(temp) / 2; j++{
				temp[j], temp[len(temp) - 1 - j] = temp[len(temp) - 1 - j], temp[j]
			}
			res = res + string(temp) + " "
			temp = []byte{}
			continue
		}else{
			temp = append(temp, s[i])
			if i == len(s) - 1{
				for j := 0; j < len(temp) / 2; j++{
					temp[j], temp[len(temp) - 1 - j] = temp[len(temp) - 1 - j], temp[j]
				}
				res = res + string(temp)
			}
		}
	}
	return
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	str := strings.Replace(input, "\n", "", -1)
	res := reverseString(str)
	fmt.Printf("%s", res)
}
