package leetcode0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	hasCommon := true
	for j := 0; hasCommon && j < len(strs[0]); j++ {
		c := strs[0][j]
		for i := 1; i < len(strs); i++ {
			if j >= len(strs[i]) || strs[i][j] != c {
				hasCommon = false
			}
		}
		if hasCommon {
			res += string(c)
		}
	}
	return res
}
