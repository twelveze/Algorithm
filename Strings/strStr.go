package main

import (
	"fmt"
	"strings"
)

func strStr(origin, aim string) int{//Sunday匹配
	ol := len(origin)
	al := len(aim)
	if ol == 0{
		return -1
	}
	if al == 0{
		return 0
	}
	if ol < al{
		return -1
	}
	originIndex := 0 //	目标串索引
	aimIndex := 0	//模式串索引
	for aimIndex < al{ //成功匹配完终止条件：所有aim均成功匹配
		if originIndex > ol - 1{// 针对origin匹配完，但aim未匹配完情况处理 如 mississippi sippia
			return -1
		}
		if origin[originIndex] == aim[aimIndex]{// 匹配则index均加1
			originIndex++
			aimIndex++
		}else{
			originIndex = originIndex - aimIndex + al // -aimIndex是为了去掉匹配到一半没匹配上的长度
			if originIndex > ol - 1{//判断下一个目标字符是否存在。
				return -1
			}else{
				step := strings.LastIndexByte(aim, byte(origin[originIndex])) // 判断目标字符在模式串中匹配到，返回最后一个匹配的index
				if step == -1{// 不存在的话，设置到目标字符的下一个元素
					originIndex = originIndex + 1
				}else{
					originIndex = originIndex - step// 存在的话，移动对应的数字（参考上文中的存在公式）
				}
				aimIndex = 0//模式串总是从第一个开始匹配
			}
		}
	}
	return originIndex - aimIndex
}
func main() {
	haystack := "Here is a Little HAO"
	handle := "Little"
	res := strStr(haystack,handle)
	fmt.Println(res)
}