package main
/*
将手中的一排格子涂上红蓝两种颜色，每个格子牛牛都有自己的想法，例如 1号和2号格子涂蓝色，
3号格子涂红色，4号格子涂蓝色，5号格子涂红色，67涂蓝色，8号涂红色。 按照这个方法，他需要6步。
但是他也可以将1-7号涂上蓝色，再将3号5号8号涂成红色,则步数最少，为4次。 但是目前数据量为500000，请你给出最少的操作次数。

输入：
8
BBRBRBBR （B表示 蓝色，R表示 红色）
输出：
4
*/
import "fmt"

func ColorGet(N int,colors string)int{
	blue := 0
	red := 0
	flag := 0
	for i := 0; i < N; i++{
		if colors[i] == 'B' && flag == 0{
			flag = 1
		}
		if flag == 1 && colors[i] == 'R'{
			flag = 0
			red++
		}
	}
	flag = 0
	for i := 0; i < N; i++{
		if colors[i] == 'R' && flag == 0{
			flag = 1
		}
		if colors[i] == 'B' && flag == 1{
			flag = 0
			blue++
		}
	}
	if colors[0] == 'R'{
		red++ //这里red代表蓝色打底最少涂red个红色
	}else{
		blue++ //这里blue表示 红色打底最少涂blue个蓝色
	}
	return min(blue, red) + 1 //红色少，就蓝色打底涂红色；蓝色少，就红色打底涂蓝色(+1 是因为要用一个颜色打底)
}
func min(a , b int) int{
	if a < b{
		return a
	}else{
		return b
	}
}
func main() {
	var N int
	fmt.Scanf("%d", &N)
	var colors string
	fmt.Scanf("%s", &colors)
	res := ColorGet(N, colors)
	fmt.Println(res)
}
