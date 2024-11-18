package main

import "fmt"

//前缀和思想
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var extraEnergy, extraExperience int
	for i := 0; i < len(energy); i++ {
		//精力上无法打败对手，需要增加精力使其满足要求
		if initialEnergy <= energy[i] {
			t := energy[i] - initialEnergy + 1
			initialEnergy += t
			extraEnergy += t
		}
		//同理，经验上无法打败对手，需要增加经验
		if initialExperience <= experience[i] {
			t := experience[i] - initialExperience + 1
			initialExperience += t
			extraExperience += t
		}
		//打败对手后更新数值
		initialEnergy -= energy[i]
		initialExperience += experience[i]
	}
	return extraEnergy + extraExperience
}

func main() {
	initialEnergy := 5
	initialExperience := 3
	energy := []int{1, 4, 3, 2}
	experience := []int{2, 6, 3, 1}
	res := minNumberOfHours(initialEnergy, initialExperience, energy, experience)
	fmt.Println(res)
}
