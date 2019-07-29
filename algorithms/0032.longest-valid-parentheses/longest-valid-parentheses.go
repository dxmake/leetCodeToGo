package leetcode0032

func longestValidParentheses(s string) int {
	// dp[i] 保存以第i个字符为结尾的有效括号串的长度
	dp := make([]int, len(s))
	max := 0
	for i := 1; i < len(s); i++ {
		// 第i个位置是有括号并且第i-1个位置为结尾的有效括号串的左边还有左括号
		if s[i] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			// 更新后的串与其左边紧邻的串拼接
			if i-dp[i] >= 0 {
				dp[i] += dp[i-dp[i]]
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}
