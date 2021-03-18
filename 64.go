package main

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[m-1][n-1] = grid[m-1][n-1]
				continue
			}
			if i+1 >= m {
				dp[i][j] = dp[i][j+1] + grid[i][j]
			} else if j+1 >= n {
				dp[i][j] = dp[i+1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
			}
		}
	}
	return dp[0][0]
}
