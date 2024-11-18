package main

/*
你在使用一个特殊的键盘输入一个字符串。键盘是一个矩形网格的形状，有各种不同的排列，每个按键上的字符互不相同，最多有 94 个按键（除空格外共有 94 个可见 ASCII 字符，ASCII 码为 33~126）。你需要操作一个机械手来点击这个键盘，机械手可以上下左右移动，每移动一格耗时 x，移动过程中（不包括一开始或者点击前后）转向耗时 y，每次点击键盘耗时 z。一开始，机械手位于左上角。请计算打完这个字符串最少需要多少时间。

输入描述
第一行五个空格隔开的整数 n, m, x, y, z，0 < n * m <= 94，1 <= x, y, z <= 100000

后面 n 行，每行一个由 m 个字符组成的字符串，表示键盘的配列，保证 m 个字符都是 ASCII 可见字符，不是空格且互不相同。

最后一行一个由上述配列中存在的字符组成的字符串，长度不超过 100000。

输出描述
一个整数，表示最少需要的时间

样例输入
2 2 1 1 1
.E
:F
EE:F.:
样例输出
15
*/
import (
	"fmt"
)

type location struct {
	x int
	y int
}

func robotWrite(targetStr string, locationMap map[byte]location, x, y, z int) int {
	var res int = 0
	lastLocation := location{0, 0}
	for _, char := range targetStr {
		location := locationMap[byte(char)]
		if location.x == lastLocation.x && location.y == lastLocation.y {
			res += z
		} else if location.x == lastLocation.x && location.y != lastLocation.y {
			res += abs(lastLocation.y, location.y) * x
			res += z
		} else if location.x != lastLocation.x && location.y == lastLocation.y {
			res += abs(lastLocation.x, location.x) * x
			res += z
		} else {
			res += abs(lastLocation.x, location.x) * x
			res += abs(lastLocation.y, location.y) * x
			res += y
			res += z
		}
		lastLocation = location
	}
	return res
}
func abs(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
func main() {
	var n, m, x, y, z int
	fmt.Scanf("%d %d %d %d %d", &n, &m, &x, &y, &z)
	locationMap := make(map[byte]location, m*n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scanln(&str)
		for j := 0; j < m; j++ {
			locationMap[str[j]] = location{i, j}
		}
	}
	var targetStr string
	fmt.Scanln(&targetStr)
	res := robotWrite(targetStr, locationMap, x, y, z)
	fmt.Println(res)
}
