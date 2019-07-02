package leetcode0001

func twoSum(nums []int, target int) []int {
	//学习使用go test，拒绝哈希
	for i := 0; ; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
}
