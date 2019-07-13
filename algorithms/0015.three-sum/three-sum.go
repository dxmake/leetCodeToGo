package leetcode0015

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	third := len(nums) - 1
	iLastNum := 0
	jLastNum := 0
	for i := 0; i < third-1; i++ {
		if i != 0 && nums[i] == iLastNum {
			continue
		}
		iLastNum = nums[i]
		for j := i + 1; j < third; j++ {
			if j != i+1 && nums[j] == jLastNum {
				continue
			}
			jLastNum = nums[j]
			for j < third && nums[i]+nums[j]+nums[third] > 0 {
				third--
			}
			if j < third && nums[i]+nums[j]+nums[third] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[third]})
			}
		}
		third = len(nums) - 1
	}
	return res
}
