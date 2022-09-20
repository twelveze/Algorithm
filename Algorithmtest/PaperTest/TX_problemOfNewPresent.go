package main

import "fmt"

func problemOfNewPresent(present string) (res int){
	for len(present) >= 3{
		for i:= 0; i < len(present); i++{
			indexa := -1
			if present[i] == '@'{
				indexa = i
			}
			if indexa != -1{
				var sum, pre, end int
				count := 1
				for i := indexa - 1;i >= 0; i--{
					if present[i] > '9' || present[i] < '0'{
						break
					}
					pre += int(present[i] - '0') * count
					count *= 10
				}
				count = 10
				end = int(present[indexa + 1] - '0')
				for i := indexa + 2; i < len(present); i++{
					if present[i] > '9' || present[i] < '0'{
						break
					}
					end = end * count
					end += int(present[i] - '0')
				}
				end = end + pre
				pre |= end
				sum = pre
				if len(present) == 3{
					return sum
				}
				sum = sum + '0'
				present = present[:indexa - 1] + string(sum) + present[indexa + 2:]
				break
			}
			indexx := -1
			if present[i] == 'x'{
				indexx = i
			}
			if indexx != -1{
				var sum, pre, end int
				count := 1
				for i := indexx - 1;i >= 0; i--{
					if present[i] > '9' || present[i] < '0'{
						break
					}
					pre += int(present[i] - '0') * count
					count *= 10
				}
				count = 10
				end = int(present[indexx + 1] - '0')
				for i := indexx + 2; i < len(present); i++{
					if present[i] > '9' || present[i] < '0'{
						break
					}
					end = end * count
					end += int(present[i] - '0')
				}
				sum = pre * end
				if len(present) == 3{
					return sum
				}
				sum = sum + '0'
				present = present[:indexx - 1] + string(sum) + present[indexx + 2:]
				break
			}
			indexj := -1
			if present[i] == '+'{
				indexj = i
			}
			if indexj != -1{
				var sum, pre, end int
				count := 1
				for i := indexj - 1;i >= 0; i--{
					if present[i] > '9' || present[i] < '0'{
						break
					}
					pre += int(present[i] - '0') * count
					count *= 10
				}
				count = 10
				end = int(present[indexj + 1] - '0')
				for i := indexj + 2; i < len(present); i++{
					if present[i] > '9' || present[i] < '0'{
						break
					}
					end = end * count
					end += int(present[i] - '0')
				}
				sum = pre + end
				if len(present) == 3{
					return sum
				}
				sum = sum + '0'
				present = present[:indexj - 1] + string(sum) + present[indexj + 2:]
				break
			}
		}
	}
	return
}
func main() {
	var present string
	fmt.Scan(&present)
	res := problemOfNewPresent(present)
	fmt.Println(res)
}
