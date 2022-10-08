package main

import (
	"fmt"
	"sort"
)

//超时(贪心双循环)
//func advantageCount(nums1, nums2 []int) []int {
//	var res []int
//	nums1Map := make(map[int]int)
//	for i := 0; i < len(nums1); i++ {
//		nums1Map[nums1[i]]++
//	}
//	sort.Ints(nums1)
//	for i := 0; i < len(nums2); i++ {
//		tempMin := math.MaxInt //记录一下当前的最小值
//		flag := true           //标志
//		for j := 0; j < len(nums1); j++ {
//			if nums1[j] < tempMin && nums1Map[nums1[j]] > 0 {
//				tempMin = nums1[j]
//			}
//			if nums1[j] > nums2[i] && nums1Map[nums1[j]] > 0 {
//				res = append(res, nums1[j])
//				nums1Map[nums1[j]]--
//				flag = false
//				break
//			}
//		}
//		//如果没有比nums2[i]大的，就把最小的nums1[j]丢进去
//		if flag {
//			res = append(res, tempMin)
//			nums1Map[tempMin]--
//		}
//	}
//	return res
//}

//思想:首先将nums1 nums2排序，然后将nums1从小到大遍历去匹配nums2,如果nums1的数字匹配不到，那么直接将其匹配nums2最大的
func advantageCount(nums1, nums2 []int) []int {
	n := len(nums1)
	idx1 := make([]int, n)
	idx2 := make([]int, n)
	for i := 0; i < n; i++ {
		idx1[i] = i
		idx2[i] = i
	}
	sort.Slice(idx1, func(i, j int) bool { return nums1[idx1[i]] < nums1[idx1[j]] })
	sort.Slice(idx2, func(i, j int) bool { return nums2[idx2[i]] < nums2[idx2[j]] })

	ans := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums1[idx1[i]] > nums2[idx2[left]] {
			ans[idx2[left]] = nums1[idx1[i]]
			left++
		} else {
			ans[idx2[right]] = nums1[idx1[i]]
			right--
		}
	}
	return ans
}

func main() {
	nums1 := []int{12, 24, 8, 32}
	nums2 := []int{13, 25, 32, 11}
	res := advantageCount(nums1, nums2)
	fmt.Println(res)
}
