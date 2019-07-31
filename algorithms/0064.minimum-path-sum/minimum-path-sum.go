package leetcode0064

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			switch {
			case i == 0 && j == 0:
				continue
			case i == 0 && j != 0:
				grid[i][j] += grid[i][j-1]
			case i != 0 && j == 0:
				grid[i][j] += grid[i-1][j]
			case i != 0 && j != 0:
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
