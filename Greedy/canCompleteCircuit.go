package main

import "fmt"

// leetcode - 134 加油站

// 该方案时间复杂度过高 --> O(n^2)
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	// 分别计算从第i个加油站开始能不能闭环
	for i := 0; i < length; i++ {
		gasArr := append(gas[i:], gas[0:i]...)
		costArr := append(cost[i:], cost[0:i]...)
		tempGas := 0
		for j := 0; j < length; j++ {
			tempGas = tempGas + gasArr[j] - costArr[j]
			if tempGas < 0 {
				break
			}
			if j == length-1 {
				return i
			}
		}
	}
	return -1
}

//去掉了无用计算 重点在44行
func canCompleteCircuit1(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		cnt, sumGas, sumCost := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumGas += gas[j]
			sumCost += cost[j]
			if sumCost > sumGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1 //假如当前一段路无法走下去了,就该放弃,换个新的起点了。这个起点最多只能到这里了，从这段路的任何地方重新开始都到达不了更远的地方了
		}
	}
	return -1
}
func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit1(gas, cost))
}
