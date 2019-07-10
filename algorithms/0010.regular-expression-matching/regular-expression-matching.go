package leetcode0010

import (
//"log"
)

func isMatch(s string, p string) bool {
	len1, len2 := len(s), len(p)
	dp := make([][]bool, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]bool, len2+1)
	}
	dp[0][0] = true
	for i := 0; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j]
			} else {
				dp[i][j] = i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') && dp[i-1][j-1]
			}
		}
	}
	// log.Printf("%v", dp)
	return dp[len1][len2]
}
