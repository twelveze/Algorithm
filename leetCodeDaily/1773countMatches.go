package main

import "fmt"

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ans := 0
	for _, iterm := range items {
		if ruleKey == "type" && ruleValue == iterm[0] {
			ans++
			continue
		}
		if ruleKey == "color" && ruleValue == iterm[1] {
			ans++
			continue
		}
		if ruleKey == "name" && ruleValue == iterm[2] {
			ans++
			continue
		}
	}
	return ans
}

func main() {
	iterms := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	res := countMatches(iterms, ruleKey, ruleValue)
	fmt.Println(res)
}
