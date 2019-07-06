package leetcode0003

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var flag [256]int
	st := 0
	end := 0
	maxLen := 0
	for {
		end++
		if flag[s[end-1]] == 0 {
			flag[s[end-1]] = 1
			if end-st > maxLen {
				maxLen = end - st
			}
		} else {
			for {
				if s[st] == s[end-1] {
					st++
					break
				} else {
					flag[s[st]] = 0
					st++
				}
			}
		}
		if end >= len(s) {
			return maxLen
		}
	}
}
