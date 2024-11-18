package main

import (
	"fmt"
	"strings"
)

// 敏感字符加密 - 华为OD 敏感字符加密
/*
	给定一个由多个命令子组成的字符串
	1、字符串的长度小于等于127字，只包含大小写字母，数字，下划线和偶数个双引号
	2、命令之间由一个或多个下划线 _ 表示
	3、可以通过两个双引号""来标识包含下划线 _ 的命令字或空命令字(仅包含两个双引号的命令字)，双引号不会在命令字内部出现

	请对指定的命令字加密，替换为 ******，并删除命令字前后多余的_
	如果无法找到指定索引的命令字 返回ERROR

	输入：输入为两行，第一行是命令字索引K(从0开始)，第二行是命令字符S
	输出：加密后的命令字或者ERROR

	示例：
	1
	password__a12345678_timeout_100
	输出
	password_timeout_100
	输入
	2
	aaa_password_"a12_45678"_timeout__100_""_
	输出
	aaa_password_timeout_100_""
*/

// 首先遍历一遍str 把包含在" "内部的_替换为$，以便后续使用 split 分割
func replace(s string) string {
	charArr := make([]byte, len(s))
	// 记录 " 的个数
	num := 0
	for i, c := range s {
		// 包含在""里的_ 才会转换为$
		if c == '_' && num > 0 {
			charArr[i] = '$'
			continue
		}
		if c == '"' && num > 0 {
			num--
		} else if c == '"' && num == 0 {
			num++
		}
		charArr[i] = byte(c)
	}
	return string(charArr)
}

//去除掉重复的_ 返回一个命令字数组
func rmRepeat_(s string) []string {
	strArr := strings.Split(s, "_")
	// 更新strArr 去除掉__ ___等_链接在一起的情况
	finalStrArr := make([]string, 0)
	for _, sub := range strArr {
		if sub == "" {
			continue
		}
		finalStrArr = append(finalStrArr, sub)
	}
	return finalStrArr
}

func encrypt(k int, s string) string {
	s = replace(s)
	strArr := rmRepeat_(s)
	fmt.Println(strArr, len(strArr))
	if k > len(strArr)-1 {
		return "ERROR"
	}
	strArr[k] = "******"
	return strings.ReplaceAll(strings.Join(strArr, "_"), "$", "_")
}
func main() {
	k := 2
	s := "aaa_password_\"a12_45678\"_timeout__100_\"\"_"
	res := encrypt(k, s)
	fmt.Println(res)
}
