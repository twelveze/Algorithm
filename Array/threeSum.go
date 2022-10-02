package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int{
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < n; i++{
		if nums[i] > 0{ //如果第一个数就大于0了，后面就不可能存在合法的三个数了
			return res
		}
		if i > 0 && nums[i] == nums[i - 1]{ //去重
			continue
		}
		left := i + 1
		right := n - 1
		for left < right{
			if nums[i] + nums[left] + nums[right] > 0{
				right--
			}else if nums[i] + nums[left] + nums[right] < 0{
				left++
			}else{
				temp := make([]int, 0)
				temp = append(temp, nums[i])
				temp = append(temp, nums[left])
				temp = append(temp, nums[right])
				res = append(res, temp)
				left++
				right--
				for nums[left] == nums[left - 1] { left++ }
				for nums[right] == nums[right + 1] { right-- }
			}
		}
	}
	return res
}
func main() {
	nums := []int{-1,0,1,2,-1,-4}
	res := threeSum(nums)
	for i := 0; i < len(res); i++{
		for j := 0; j < len(res[i]); j++{
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println()
	}
}
