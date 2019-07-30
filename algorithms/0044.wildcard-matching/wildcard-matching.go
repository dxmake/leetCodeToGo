package leetcode0044

func isMatch(s string, p string) bool {
	// dp[i][j]存储s[0:i]和p[0:j]匹配成功
	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	// dpStar[j]在第i层循环时表示s[0:k]和p[0:j]匹配成功,k<=i
	dpStar := make([]bool, len(p))
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			dpStar[j-1] = dpStar[j-1] || dp[i][j-1]
			switch p[j-1] {
			case '*':
				dp[i][j] = dpStar[j-1]
			case '?':
				if i != 0 {
					dp[i][j] = dp[i-1][j-1]
				}
			default:
				if i != 0 {
					dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
