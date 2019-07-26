package leetcode0188

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxProfit(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	k = min(k, len(prices)/2)
	dp := make([]int, len(prices))
	tmp := 0
	newTmp := 0
	for i := 1; i <= k; i++ {
		tmp = -prices[0]
		for j := 1; j < len(prices); j++ {
			newTmp = max(tmp, dp[j]-prices[j])
			dp[j] = max(dp[j-1], tmp+prices[j])
			tmp = newTmp
		}
	}
	return dp[len(prices)-1]
}
