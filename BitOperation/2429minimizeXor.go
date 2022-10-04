package main

import (
	"fmt"
	"math/bits"
)

//x的置位数与num2相同,记为c2,所以我们要合理分配这c2个1
//为了让异或和尽量小，这些1应当从高位到低位匹配，如果匹配完了还有多余的，那么就从低位到高位把0改成1
//考虑直接在num1上修改，分类讨论:
//如果c2 < c1:直接把num1最低位的c1-c2个1变成0
//如果c2 > c1:直接把num1最低位的c2-c1个0变成1
func minimizeXor(num1 int, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))

	for ; c2 < c1; c2++ {
		num1 &= num1 - 1 // 最低的 1 变成 0
	}

	for ; c2 > c1; c2-- {
		num1 |= num1 + 1 // 最低的 0 变成 1
	}
	return num1
}

func minimizeXor1(num1 int, num2 int) int {
	bitMap := make(map[int]bool, 30) //这里设为30并没有什么基本的含义，因为2^30肯定大于题目要求上限10^9
	count := 0
	//求出num2 1的个数
	for num2 > 0 {
		if num2&1 == 1 {
			count++
		}
		num2 = num2 / 2
	}
	ans := 0
	//先从高位抵消
	for i := 30; i >= 0; i-- {
		if (num1>>i)&1 == 1 {
			bitMap[i] = true
		} else {
			bitMap[i] = false
		}
		if bitMap[i] && count > 0 { //这一位是1(跟num1的高位异或抵消)
			count--
			ans += 1 << i
		}
	}
	//如果count多余，从低位开始构建,要尽可能的使1在低位,这样异或得出来的数字可以尽量小
	for i := 0; i < 30 && count > 0; i++ {
		if bitMap[i] { //num1低位i为1,这时ans的i位只能为0
			continue
		} else { //反之
			count--
			ans += i << i
		}
	}
	return ans
}

func main() {
	num1, num2 := 3, 5
	res := minimizeXor(num1, num2)
	res1 := minimizeXor1(num1, num2)
	fmt.Println("res: ", res)
	fmt.Println("res1: ", res1)
}
